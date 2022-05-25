package customR

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/testWay", testWayHandler)
	testGetParams(r)
	testGroupRouper(r)
	// 测试分模块加载路由功能
	LoadBlogRouter(r)
	LoadForumRouter(r)

	return r
}

func testWayHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"message": "no message",
		})
}

func testGetParams(r *gin.Engine) {
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.GET("/getUserInfo", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default")
		c.String(http.StatusOK, "hello."+name)
	})
}

/// 测试组路由, 挺有用的功能
func testGroupRouper(r *gin.Engine) {
	v1 := r.Group("/V1")
	v1.POST("/login", loginV1)
	v1.POST("/submit", submitV1)

	v2 := r.Group("/V2")
	v2.POST("/login", loginV2)
	v2.POST("/submit", submitV2)

	r.GET("/getTestJson", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"message": "this is test json",
		})
}

func loginV1(c *gin.Context) {
	c.String(http.StatusOK, "loginV1")
}
func submitV1(c *gin.Context) {
	c.String(http.StatusOK, "submitV1")
}
func loginV2(c *gin.Context) {
	c.String(http.StatusOK, "loginV2")
}
func submitV2(c *gin.Context) {
	c.String(http.StatusOK, "submitV2")
}
