package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func Test_Cought(t *testing.T) {
	cleanIps()
	Cought()
	fi, err := ioutil.ReadDir("ips")
	if err != nil {
		log.Println("ips文件夹读取失败")
		log.Fatalln(err)
	}
	if len(fi) > 0 {

		log.Println("ip.txt抓取测试通过", fi[0].Name())
	} else {
		log.Println("ip.txt抓取测试失败")

	}

}
