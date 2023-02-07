//控制器代码
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"proxy-website/dao"
	"proxy-website/tools"
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

func loginHtml(c *gin.Context) {
	msg := c.Query("msg")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": msg,
	})
}

/*
	网页登录
*/
func loginForm(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	c.ShouldBind(&json)
	username := json["username"]
	passowrd := json["password"]
	passowrd = tools.PasswordEncode(passowrd)
	user := dao.FindUser(username, passowrd)
	if user.ID == 0 {
		c.JSON(http.StatusOK, map[string]string{"code": "0", "msg": "用户名或密码错误"})
	} else {
		token, err := tools.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"code": "-1", "msg": "未知错误"})
			return
		}
		// c.SetCookie("cookieName", "name", 10, "/", "", true, true)
		c.SetCookie("token", token, 0, "/", "", false, false)

		c.JSON(http.StatusOK, map[string]string{"code": "1", "msg": "登陆成功"})

	}
}

/*
	获取ip.txt版本，检查是否需要更新
*/
func ipTest(c *gin.Context) {
	c.JSON(http.StatusOK, dao.FindIpFirst())
}

func ipGet(c *gin.Context) {
	fi, err := ioutil.ReadDir(SAVE_PATH)
	data := map[string]interface{}{}
	if err != nil {
		data["msg"] = "文件夹打开异常"
		c.AbortWithStatusJSON(http.StatusInternalServerError, data)
	}
	if len(fi) == 0 {
		data["msg"] = "未找到ip文件"
		c.AbortWithStatusJSON(http.StatusInternalServerError, data)
	}
	c.File("static/ip/" + fi[0].Name())
}

type VC struct {
	code     string    // 验证码
	deadline time.Time //截止日期
}

var VerifyCodes = make(map[string]VC)

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

/*
	每15分钟清理一次超时且未使用的验证码
*/
func init() {
	go func() {
		for ; ; time.Sleep(time.Minute * 15) {
			for key, val := range VerifyCodes {
				if !val.deadline.Before(time.Now()) {
					delete(VerifyCodes, key)
				}
			}
		}
	}()
}

/*
	发送验证码
*/
func verifycode(c *gin.Context) {
	json := make(map[string]string)
	c.ShouldBind(&json)
	if username, ok := json["username"]; ok {
		code := AuthCode(6)
		SendAuthCodeByMail(code, username)
		VerifyCodes[username] = VC{code: code, deadline: time.Now().Add(time.Minute * 30)} //验证码存在半小时
		c.JSON(http.StatusOK, map[string]string{"msg": "验证码已经发送到邮箱"})
	} else {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "后台未收到用户名"})
	}
}

/*
	用户注册
*/

func registerForm(c *gin.Context) {
	json := make(map[string]string)
	c.ShouldBind(&json)

	username := json["username"]
	passowrd := json["password"]
	verifycode := json["verifycode"]
	if username == "" || passowrd == "" || verifycode == "" {
		c.Redirect(http.StatusFound, "/register?msg=缺少请求参数")
	}

	user := dao.FindUserByUserName(username)
	if user.ID != 0 {
		c.JSON(http.StatusOK, map[string]string{"code": "0", "msg": "账号已经被注册，修改密码请联系管理员"})
		return
	}

	if VerifyCodes[username].code != verifycode {
		c.JSON(http.StatusOK, map[string]string{"code": "0", "msg": "验证码错误"})
		return
	}
	user = &dao.User{UserName: username, Pwd: passowrd}
	err := dao.UserRegister(user)
	if err != nil {
		c.JSON(http.StatusFound, map[string]string{"code": "0", "msg": "未知错误：" + fmt.Sprintln(err)})
	} else {
		delete(VerifyCodes, username)
		c.JSON(http.StatusOK, map[string]string{"code": "1", "msg": "注册成功"})
	}

}

func bind(c *gin.Context) {
	json := make(map[string]string)
	c.ShouldBind(&json)

	if key, exists := json["key"]; exists {
		username := c.GetString("UserName")
		dao.BindKey(key, username)
		c.JSON(http.StatusOK, map[string]interface{}{"code": "1", "msg": "绑定成功"})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"code": "0", "msg": "未接收到密钥"})
	}
}

func mainUserPage(c *gin.Context) {
	s := c.GetString("UserName")
	result, _ := dao.FindTrafficByUser(s)

	c.HTML(http.StatusOK, "main.html", map[string]interface{}{"traffic": result})
}
