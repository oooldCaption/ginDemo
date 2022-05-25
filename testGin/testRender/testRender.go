package testRender

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

/// 测试文件格式渲染
func TestRender() {
	r := gin.Default()

	testJSONRender(r)
	testStructRender(r)
	testXMLRender(r)
	testYAMLRender(r)
	testProtobufRender(r)

	r.Run()
}

/// 测试 json 渲染
func testJSONRender(r *gin.Engine) {
	r.GET("/testJSONRender", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "JSON DATA"})
	})
}

/// 测试 struct 渲染
func testStructRender(r *gin.Engine) {
	r.GET("/testStruct", func(context *gin.Context) {
		var msg struct {
			Name    string
			Message string
			NUmber  int
		}
		msg.NUmber = 23
		msg.Message = "this is struct"
		msg.Name = "煎人寿"
		context.JSON(200, msg)
	})
}

/// 测试 xml 渲染
func testXMLRender(r *gin.Engine) {
	r.GET("/testXMLRender", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "XML DATA"})
	})
}

func testYAMLRender(r *gin.Engine) {
	r.GET("/testYAMLRender", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "YAML DATA"})
	})
}

func testProtobufRender(r *gin.Engine) {
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
}
