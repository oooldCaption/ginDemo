package reviewGo

import "fmt"

func RunDay11() {
	libDemo()
}

//模拟终端图书管理系统

// Book 图书结构
type Book struct {
	ID     int
	Title  string
	Author string
}

func libDemo() {
	fmt.Println("欢迎使用图书管理系统")
	println("1. 添加图书")
	println("2. 删除图书")
	println("3. 列出所有图书")
	println("4. 查找图书")
	println("5. 退出")
	checkCmd()
}
func checkCmd() {
	var cmd int
	_, err := fmt.Scanln(&cmd)
	if err != nil {
		println("输入错误")
		return
	}
	switch cmd {
	case 1:
		addBook()
	case 2:
		delBook()
	case 3:
		listBook()
	case 4:
		searchBook()
	case 5:
		exit()
	default:
		exit()
	}

}

var bookList []Book
var currentID int

func addBook() {
	var title string
	var author string

	println("请输入 书名")
	fmt.Scanln(&title)
	println("请输入 作者")
	fmt.Scanln(&author)

	b := Book{
		ID:     currentID,
		Title:  title,
		Author: author,
	}
	currentID += 1
	bookList = append(bookList, b)
	endLogic()

}

func delBook() {
	var bookId int
	println("请输入要删除的书本 ID")
	fmt.Scanln(&bookId)
	var idx int
	for i, book := range bookList {
		if bookId == book.ID {
			idx = i
		}
	}
	bookList = append(bookList[:idx], bookList[idx+1:]...)
	listBook()

	endLogic()
}

func listBook() {
	for _, book := range bookList {
		println("当前读书列表:", book.Title, book.Author)
	}
	println()
	println()
	endLogic()
}

func searchBook() {
	println("请输入要搜索的内容:")
	var info int
	fmt.Scanln(&info)
	for _, book := range bookList {
		if info == book.ID {
			println("搜索结果:", book.Title, book.Author)
		}
	}

	endLogic()

}

func exit() {

}

func endLogic() {
	libDemo()
}

func day11Ex1() {

	var s []string
	s = append(s, "1")
	fmt.Println(s)

	// panic: assignment to entry in nil map
	// 空 map 不能这样赋值
	var m map[string]string
	// 使用 make 进行初始化
	//var m = make(map[string]string)
	m["name"] = "qk"
	m["age"] = "23"
	fmt.Println(m)
}
