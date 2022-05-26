package testCookie

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func TESTSaveSession() {
	r := gin.Default()
	r.Use(func(context *gin.Context) {
		sessions.NewSession(store, "mySession")
	})
	r.GET("/saveSession", saveSession)
	r.GET("/getSession", getSession)
	r.Run()
}

func saveSession(c *gin.Context) {
	s, e := store.Get(c.Request, "mySession")
	if e != nil {
		c.Error(e)
	}
	s.Values["name"] = "煎人寿"
	s.Values["age"] = "29"
	s.Save(c.Request, c.Writer)
}

func getSession(c *gin.Context) {
	s, e := store.Get(c.Request, "mySession")
	if e != nil {
		c.Error(e)
	}

	name := s.Values["name"]
	age := s.Values["age"]
	fmt.Println(name, age)
	//c.JSON(200,{"name":name,"age":age})
}
