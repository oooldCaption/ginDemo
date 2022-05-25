package customR

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadForumRouter(r *gin.Engine) {
	r.GET("/forumOne", forumHandler)
	r.GET("/forumTwo", forumHandler)
}

func forumHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is Forum message",
	})
}
