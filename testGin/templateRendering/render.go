package templateRendering

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// 设置模板文件的路径
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		// 渲染模板
		c.HTML(200, "index.html", gin.H{
			"Name": "Gin User",
		})
	})

	r.GET("/map", func(c *gin.Context) {

		c.HTML(200, "map.html", gin.H{
			"title":       "Hello from Map",
			"description": "This is a description using map.",
		})

	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://52smile.vip")
	})

	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	r.Run(":8080")
}
