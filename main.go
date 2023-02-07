package main

import (
	"html/template"
	"net/http"
	"proxy-website/env"
	"proxy-website/tools"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		token, err2 := tools.ParseToken(s)
		if err2 != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if mc, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("UserName", mc["UserName"])
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

	}
}

func formatAsDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/ip/test", ipTest)
		authorized.GET("ip/get", ipGet)
		authorized.GET("/main", mainUserPage)
		authorized.POST("/bind", bind)
	}

	r.LoadHTMLGlob("static/template/*")
	r.GET("/", indexHtml)
	r.GET("/index", indexHtml)
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/release/", "./static/release/")
	r.Static("/static/", "./static/static/")
	r.GET("login", loginHtml)
	r.POST("login/form", loginForm)
	r.GET("register", register)
	r.POST("register/form", registerForm)
	r.POST("/verifycode", verifycode)
	_ = r.Run(":" + strconv.Itoa(env.GetConfig().Port))
}
