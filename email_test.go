package main

import (
	"fmt"
	"testing"
)

func Test_SendToMail(t *testing.T) {

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

	fmt.Printf(body, "12345")

}

func Test_RandInt(t *testing.T) {

	dir := make(map[string]int)
	for i := 0; i < 10; i++ {
		dir[AuthCode(6)]++
	}

	for key, val := range dir {
		fmt.Println(key, ",", val)
	}
}
