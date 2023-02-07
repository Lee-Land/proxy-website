package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

/*
账号 liu.proxy.qq.com
密码 krexqcjubwjtdieg
*/
func SendAuthCodeByMail(code, to string) {
	user := "liu.proxy@qq.com"
	password := "krexqcjubwjtdieg"
	host := "smtp.qq.com:25"

	subject := "LiuProxy验证码"

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			<h2>验证码是【%s】</h2>
		</body>
		</html>`
	body = fmt.Sprintf(body, code)
	sendUserName := "LiuProxy" //发送邮件的人名称
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func AuthCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	time.Sleep(time.Nanosecond)
	return sb.String()
}
