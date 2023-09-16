package testDataFmt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `json:"user" binding:"required" uri:"user" xml:"user" form:"user"`
	PSW  string `json:"psw"  binding:"required" uri:"psw" xml:"psw" form:"psw"`
}

func TestDataFMT() {
	r := gin.Default()
	testURI(r)
	r.Run()
}

func TestJSON(r *gin.Engine) {
	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if json.User != "root" || json.PSW != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "login success"})

	})
}

func TestForm(r *gin.Engine) {
	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		// bind 默认解析 form 表单. 根据请求头中的content-type自动推断
		if err := c.Bind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if json.User != "root" || json.PSW != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "login success"})

	})
}

func testURI(r *gin.Engine) {
	r.POST("/blog/category/:user/tag/:psw", func(c *gin.Context) {
		var json Login
		// bind 默认解析 form 表单. 根据请求头中的content-type自动推断
		if err := c.BindUri(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": json.User,
			"psw":  json.PSW,
		})
		//if json.User != "root" || json.PSW != "admin" {
		//	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		//	return
		//}
		//c.JSON(http.StatusOK, gin.H{"status": "login success"})

	})

}
