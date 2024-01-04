package sever

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"log"
	"sort"
	"strings"
	"time"
	"wxlogin/cfg"
	"wxlogin/dao"
	"wxlogin/models"
	//"wxlogin/sever"
)

//const Token = "coleliedev"

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// Sha1 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func WXCheckSignature(c *fiber.Ctx) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := CheckSignature(signature, timestamp, nonce, cfg.Config.Sever.WxToken)
	if !ok {
		log.Println("[微信接入] - 微信公众号接入校验失败!")
		return
	}

	log.Println("[微信接入] - 微信公众号接入校验成功!")
	//_, _ = c.Writer.WriteString(echostr)
	_ = c.SendString(echostr)

}

// WxTextMsg 微信文本消息结构体 正常使用
func WxTextMsg(c *fiber.Ctx) {
	var textMsg models.WXTextMsg
	if err := xml.Unmarshal(c.Body(), &textMsg); err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n\n", textMsg.MsgType, textMsg.Content)
	if textMsg.Content != "领取" {
		WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName, cfg.Config.App.OtherSub)
		return
	}
	d, err := dao.RedisClient.Get(dao.Ctx, textMsg.FromUserName).Result()
	log.Print(d, err)
	if err != redis.Nil && len(d) != 0 {
		log.Printf("[消息发送] - FromUserName: %s 已经发送过消息，等待24小时后再次发送\n", textMsg.FromUserName)
		WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName, cfg.Config.App.UsedSub)
		return
	}
	k, err := SetToken()
	if err != nil {
		log.Print("获取失败：", err)
		return
	}
	WXMsgReply(c, textMsg.ToUserName, textMsg.FromUserName, fmt.Sprintf(cfg.Config.App.TokenSub, k))
	err = dao.RedisClient.Set(dao.Ctx, textMsg.FromUserName, "1", 24*time.Hour).Err()
	if err != nil {
		log.Printf("[消息发送] - FromUserName: %s 发送消息失败\n", textMsg.FromUserName)
		return
	}
}

func FirstWxMsg(c *fiber.Ctx) {
	var msg models.WeChatMessage
	if err := xml.Unmarshal(c.Body(), &msg); err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n\n", msg.MsgType, msg.Event)
	if msg.Event != "subscribe" {
		return
	}
	repTextMsg := models.WXRepTextMsg{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      cfg.Config.App.FirstSub,
	}
	data, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_ = c.Send(data)
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *fiber.Ctx) {
	var msgType models.WeChatMsgType
	if err := xml.Unmarshal(c.Body(), &msgType); err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	log.Printf("[消息接收] - 收到消息, 消息类型为: %s\n", msgType.MsgType)
	if msgType.MsgType == "text" {
		WxTextMsg(c)
	} else if msgType.MsgType == "event" {
		FirstWxMsg(c)
	}
}

// WXMsgReply 微信消息回复
func WXMsgReply(c *fiber.Ctx, fromUser, toUser, m string) {
	repTextMsg := models.WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      m,
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_ = c.Send(msg)
}
