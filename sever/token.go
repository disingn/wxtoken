package sever

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"wxlogin/cfg"
	"wxlogin/models"
)

func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.NewSource(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func SetToken() (string, error) {
	url := cfg.Config.Sever.APIUrl + "api/token/"
	method := "POST"
	id := generateRandomString(10)

	data := models.SetToken{
		Name:           id,
		RemainQuota:    cfg.Config.Sever.Limit,
		ExpiredTime:    time.Now().Add(time.Hour * 24 * time.Duration(cfg.Config.Sever.Expire)).Unix(),
		UnlimitedQuota: false,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Print("data json 转换失败")
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", cfg.Config.Sever.Authorization)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Host", "www.bxsapi.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "X-ANTS-WAF-R-C=0001664324")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Print("Error:", res.StatusCode)
		return "", fmt.Errorf("error:%d", res.StatusCode)
	}
	k, err := GetToken(id)
	if err != nil {
		log.Print("获取token失败")
		return "", err
	}
	return "sk-" + k, nil
}

func GetToken(id string) (string, error) {
	url := cfg.Config.Sever.APIUrl + "api/token/?p=0&size=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", cfg.Config.Sever.Authorization)
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Host", "www.bxsapi.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "X-ANTS-WAF-R-C=0001664324")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if res.StatusCode != 200 {
		fmt.Println("Error:", string(body))
		return "", err
	}
	var data models.GetToken
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return "", err
	}
	if data.Data[0].Name != id {
		return "", fmt.Errorf("获取失败")
	}
	return data.Data[0].Key, nil
}
