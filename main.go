package main

import (
	"net/http"
	"proxy-website/env"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome to Liu-Proxy Website",
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
	_ = r.Run(":" + strconv.Itoa(env.GetConfig().Port))
}
