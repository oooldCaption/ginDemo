package testCookie

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestCookie() {

	r := gin.Default()
	//testGetCookie(r)
	testAuthCookie(r)
	r.Run()
}

/// how to get cookie
func testGetCookie(r *gin.Engine) {
	r.GET("/testCookie", func(context *gin.Context) {
		cookie, err := context.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			context.SetCookie("key_cookie", "value_cookie", 3600, "/", "localhost", false, true)
		}
		context.JSON(200, cookie)
		fmt.Println(cookie)
	})
}

/*

注意 domain localhost != 127.0.0.1
cookie 传输不安全 , 因为是明文传递
增加带宽消耗
可以被禁用
cookie长度存在上限

*/
func testAuthCookie(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.JSON(200, "Login Success")
	})
	r.GET("/home", authMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}
