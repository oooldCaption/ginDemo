package customR

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlogRouter(r *gin.Engine) {
	r.GET("/blogOne", blogHandler)
	r.GET("/blogTwo", blogHandler)
}

func blogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is Blog message",
	})
}
