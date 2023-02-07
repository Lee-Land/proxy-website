package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func Test_Cought(t *testing.T) {
	cleanIps()
	Cought()
	fi, err := ioutil.ReadDir(SAVE_PATH)
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

func TestXxx(t *testing.T) {
	has := md5.Sum([]byte("123456"))
	pwd := fmt.Sprintf("%x", has)
	fmt.Println(pwd)
}
