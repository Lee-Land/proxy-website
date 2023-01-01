//爬虫。爬取国内ip范围，保存到ip.txt
package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	ADDRESS   = "http://ip.bczs.net/country/CN"
	SAVE_PATH = "static/ip"
)

func queryDate(headText string) string {
	match := regexp.MustCompile(`20[\d]{2}-[\d]{2}-[\d]{2}`)
	s := match.FindStringSubmatch(headText)
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

func cleanIps() {
	os.RemoveAll(SAVE_PATH)
	os.Mkdir(SAVE_PATH, 0666)
}

func Cought() {
	r, err := http.Get(ADDRESS)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Panicln(err)
	}

	headText := doc.Find("thead").Find("td").Text()

	f, err := os.Create(SAVE_PATH + "/ip_" + queryDate(headText) + ".txt")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		f.WriteString(s.Find("td:nth-child(1)").Text() + " " + s.Find("td:nth-child(2)").Text() + "\n")
	})

}

func init() {
	log.Println("国内ip范围爬取服务启动")
	go func() {
		cleanIps()
		Cought()
		time.Sleep(time.Hour * 24)
	}()
}
