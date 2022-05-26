package testMiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/// 测试中间件

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("开始执行中间件")
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println(t2)
	}
}

func TestMiddleWare() {
	r := gin.Default()
	r.Use(MiddleWare())
	{
		r.GET("/testMiddleWare", func(context *gin.Context) {
			req, _ := context.Get("request")
			context.JSON(200, gin.H{"request": req})
		})
	}
	r.Run()
}

/// 除了可以防止全局中间件, 还可以放置 局部中间件
