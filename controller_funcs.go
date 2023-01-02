//控制器代码
package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ipTest(c *gin.Context) {
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
	data["name"] = fi[0].Name()
	c.JSON(http.StatusOK, data)
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
