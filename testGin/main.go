package main

import (
	"fmt"
	"ginDemo/testGin/customR"
	"ginDemo/testGin/testOther"
	"github.com/gin-gonic/gin"
)

func main() {
	//testFileRouter()
	//testDataFmt.TestDataFMT()
	//testRender.TestRender()
	//testMiddleWare.TestMiddleWare()
	//testCookie.TestCookie()
	//testCookie.TESTSaveSession()
	testOther.TestLogger()
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
