package language

import (
	"fmt"
)

// Books type is a
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

// OopTest is to try golang OOP
func OopTest() {
	var book1 Books
	book1.title = "aabbcc"
	book1.author = "zhufeng"
	book1.bookID = 1
	book1.subject = "salary=1000000"

	var book2 Books
	book2.title = "不对"
	book2.author = "朱峰"
	book2.bookID = 2
	book2.subject = "ruhe "
	fmt.Println(book1)
	fmt.Println(hehe(book1, book2))
}

func hehe(b1, b2 Books) string {

	return b1.title + " ..." + b2.title
}

func OopFuncTests() {
	book := Books{}
	fmt.Printf("book struct address:%p\n", &book)
	book.SetInfo("Go语言", "朱峰", 101, "Go语言入门到放弃")
	book.printInfoByParameter()
	book.printInfoByPointer()
}

func (b Books) printInfoByParameter() {
	fmt.Printf("lets print first for variable b struct address: %p\n", &b)
	fmt.Printf("then print info: title:%s, author:%s, subject:%s, bookID:%d\n", b.title, b.author, b.subject, b.bookID)
}

func (b *Books) printInfoByPointer() {
	fmt.Printf("lets print first  for variable b pointer address:%p\n", b)
	fmt.Printf("then print info: title:%s, author:%s, subject:%s, bookID:%d\n", b.title, b.author, b.subject, b.bookID)
}

func (b *Books) SetInfo(title, author string, bookID int, subject string) {
	b.title = title
	b.author = author
	b.bookID = bookID
	b.subject = subject
}
