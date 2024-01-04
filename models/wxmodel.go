package models

import "encoding/xml"

type WXTextMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId,omitempty"`
	Idx          string   `xml:"Idx,omitempty"`
}

type WeChatMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   // 目标用户的用户名
	FromUserName string   `xml:"FromUserName"` // 发送方的用户名
	CreateTime   int64    `xml:"CreateTime"`   // 消息创建时间（整型）
	MsgType      string   `xml:"MsgType"`      // 消息类型，这里是 "event"
	Event        string   `xml:"Event"`        // 事件类型，这里是 "subscribe"
}

type WeChatMsgType struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string   `xml:"MsgType"` // 消息类型，这里是 "event"
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}
