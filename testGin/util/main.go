package util

import (
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type GoUser struct {
	Id   string `orm:"id,primary_key" json:"id"`
	Name string `orm:"name" json:"name"`
	Age  int    `orm:"age" json:"age"`
}

/***********************测试数据库功能********************************/
var Db *sqlx.DB

func Init() {

	database, err := sqlx.Open("mysql", "root:Suqiankun9@tcp(192.168.50.222:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	Db = database

}

func Save(c *gin.Context) {
	fmt.Println("save func call")
	var user GoUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request, err := Db.NamedExecContext(c, "INSERT INTO go_user (id,name,age) VALUES (:id,:name,:age)", &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Update() {

}

type User struct {
	Name     string `db:"name"`
	Psw      string `db:"psw"`
	NickName string `db:"nickName"`
	Age      int    `db:"age"`
	Sign     string `db:"sign"`
	UserID   int    `db:"userID"`
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Psw, u.NickName, u.Age, u.Sign}, nil
}

func Get() {
	sql := "select * from user where userID > ?"
	var users []User
	err := Db.Select(&users, sql, 1)
	fmt.Println("11111")
	if err != nil {
		fmt.Println("2222")
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}

func InsertUser() {
	_, err := Db.NamedExec(`insert into user (name,psw,nickName,age,sign) values (:name,:psw,:nickName,:age,:sign)`,
		User{
			Name:     "ali",
			Psw:      "12345",
			NickName: "阿里",
			Age:      32,
			Sign:     "阿里最喜欢跳舞兰",
		},
	)
	if err != nil {
		fmt.Println("插入失败:", err)
		return
	}

}

func SendSome() {
	u1 := User{
		Name:     "xiaolv",
		Psw:      "12345",
		NickName: "陆婉秋",
		Age:      23,
		Sign:     "错错错",
	}
	u2 := User{
		Name:     "xiaomeng",
		Psw:      "12345",
		NickName: "孟浩然",
		Age:      23,
		Sign:     "天涯辣个不识君",
	}
	u3 := User{
		Name:     "xiaoli",
		Psw:      "12345",
		NickName: "李白",
		Age:      23,
		Sign:     "忽而将换美酒出",
	}
	u4 := User{
		Name:     "xiaoli",
		Psw:      "12345",
		NickName: "李白",
		Age:      23,
		Sign:     "忽而将换美酒出",
	}

	sql, args, err := sqlx.In(
		`insert into user (name,psw,nickName,age,sign) values `+`(?),(?),(?)`,
		[]interface{}{u1, u2, u3, u4}...,
	)
	if err != nil {
		fmt.Println("处理错误:", err)
		return
	}

	fmt.Println(sql)
	fmt.Println(args)
	Db.Exec(sql, args...)

}
