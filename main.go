package main

import (
	"fmt"
	"ginDemo/testGin/customR"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//testFileRouter( )
	testViper()
	//testDataFmt.TestJSON()
}

func testViper() {
	viper.SetDefault("filePath", "./")

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

type tester interface {
	delRoot()
	addRoot()
}
type test struct {
}

func sourceCode() {
	t := test{}
	if _, ok := interface{}(t).(tester); ok {

	} else {

	}
}
