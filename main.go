package main

import (
	"io/ioutil"
	"net/http"
	"proxy-website/env"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHtml(c *gin.Context) {
	files, err := ioutil.ReadDir("./static/release/")
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	var releaseFiles []string
	for _, file := range files {
		releaseFiles = append(releaseFiles, file.Name())
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome to Liu-Proxy Website",
		"files": releaseFiles,
	})
}

func main() {
	r := gin.Default()
	r.GET("/checkup", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now().String())
	})

	r.LoadHTMLGlob("static/template/*")
	r.GET("/", indexHtml)
	r.GET("/index", indexHtml)
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/release/", "./static/release/")
	_ = r.Run(":" + strconv.Itoa(env.GetConfig().Port))
}
