package main

import (
	"context"
	"fmt"
	"ginDemo/testGin/customR"
	"ginDemo/testGin/noSql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	//testFileRouter()
	//testDataFmt.TestDataFMT()
	//testRender.TestRender()
	//testMiddleWare.TestMiddleWare()
	//testCookie.TestCookie()
	//testCookie.TESTSaveSession()
	//testOther.TestLogger()
	//hash := md5.New()
	//hash.Write([]byte("SuQiankun"))
	//hashInBytes := hash.Sum(nil)
	//hashInStr := hex.EncodeToString(hashInBytes)
	//fmt.Println(hashInStr)
	r := Router()
	r.Run(":8989")
}

/***********************测试路由功能********************************/
/// 测试抽取单个路由文件
type Option func(engine *gin.Engine)

var opts = []Option{}

func IncloudRouter(o ...Option) {
	opts = append(opts, o...)
}
func testFileRouter() {
	/// 还可以通过这种方式定义一个按需加载路由功能的方法
	//IncloudRouter(customR.LoadForumRouter, customR.LoadBlogRouter)
	r := customR.InitRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err %v\n", err)
	}
}

/***********************数据库添加api接口********************************/

func Router() *gin.Engine {

	if err := noSql.InitRedis(); err != nil {
		fmt.Println("初始化数据库错误:", err)
	}

	//err := noSql.WriteKeyWithName(context.Background(), "name", "value")
	//noSql.GetValueWithKey("dsa", context.Background())
	noSql.GetZSet(context.Background())

	r := gin.Default()
	s := r.Group("/user")
	s.GET("/save", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user": "su",
			"psw":  "qiankun",
		})
	})
	//s.GET("/select", util.Get)

	return r
}
