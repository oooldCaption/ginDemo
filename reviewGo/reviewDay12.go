package reviewGo

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"unicode/utf8"
)

func RunDay12() {
	//creatTable()
	day12Ex3()
}

func day12Ex3() {
	//a := [3]int{1, 2, 3}
	//println(len(a))
	//println(cap(a))

	s := make([]int, 0, 0)
	s = append(s, 1, 2, 3)

	ss := s[1:2]
	// 增加后 修改了 ss 的结构, 这时候不与 s 共享结构
	//ss = append(s, 1099)
	ss[0] = 99

	for _, i2 := range s {
		println(i2)
	}

	//println(len(s))
	//println(cap(s))
	//println(s[:1])
	//println(s[1:2])

}

func simpleWeb() {
	// 官方 demo;
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", req.URL.Path[1:])
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func day12Ex1() {
	//s := "this is way"
	s := "😚你好"
	println(len(s))
	println(utf8.RuneCountInString(s))

}

func day12Ex2() {
	t := 3.14242132132132
	ratio := math.Pow(10, t)
	println(math.Round(t*ratio) / ratio)
}

// 书店管理系统
// 数据模型:
// 1. 书籍信息
type BookInfo struct {
	//gorm.Model
	Title  string
	ID     int
	Author Author `gorm:"embedded"`
	Price  float64
}

// 2. 作者信息
type Author struct {
	gorm.Model
	Name      string
	BirthYear int
}

// 3.销售信息
type SaleRecord struct {
	//gorm.Model
	Book   BookInfo `gorm:"embedded"`
	Amount int
	Date   string
}

// 4. 创建数据表
var db *gorm.DB

func creatTable() {
	dsn := "root:Suqiankun9@tcp(192.168.50.222:3306)/gin_blog?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = d
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&BookInfo{})
	db.AutoMigrate(&SaleRecord{})
}

// 功能列表:
// 0. 主菜单
func mainMenu() {
	println("欢迎使用书店管理系统!")
	var cmd int
	fmt.Scanln(&cmd)
	switch cmd {
	case 1:
		day12AddBook()
	case 2:
		day12DelBook()
	case 3:
		day12SearchBook()
	case 4:
		day12AddAuthor()
	case 5:
		day12ListBook()
	case 6:
		day12SaleBook()
	case 7:
		day12ListSaleRec()
	case 8:
		day12ListSaleAmount()
	case 9:
		day12Exit()
	}
}

// 1. 添加书籍
func day12AddBook() {
	println("请输入书本名")
	var bookName string
	fmt.Scanln(&bookName)
}

// 2. 删除书籍
func day12DelBook() {

}

// 3. 查找书籍
func day12SearchBook() {

}

// 4. 添加作者
func day12AddAuthor() {

}

// 5. 查看所有书籍
func day12ListBook() {

}

// 6. 销售书籍
func day12SaleBook() {

}

// 7. 查看销售记录
func day12ListSaleRec() {

}

// 8. 查看销售总额
func day12ListSaleAmount() {

}

// 9. 退出
func day12Exit() {
	mainMenu()
}
