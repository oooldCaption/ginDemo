package testOther

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

/// 测试写入 日志文件
func TestLogger() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//gin.DefaultWriter = io.MultiWriter()
	r := gin.Default()
	r.GET("/pingLog", func(context *gin.Context) {
		context.JSON(200, "pingLog")
	})
	r.Run()

}
