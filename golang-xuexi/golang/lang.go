package golang

import "fmt"

// Books type is a
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

// LangTest is package main
func LangTest() {
	fmt.Println("go language study")
	//pointer()
	//struct_and_string()
	//array()
	//defer_call()
}

func TypeTest(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Printf("value is string = %s\n", str)
	case int:
		fmt.Printf("value is int = %d\n", str)
	}
}

func defer1() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func defer2(x, y int) (z int) {
	defer func() { z += 100 }()
	z = x + y
	return z + 50 // 执行顺序 z = z+50 -> (call defer)z = z+100 -> ret
}

func defer3() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func defer4() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func deferTest() {
	defer1()
	fmt.Println()
	fmt.Println(defer2(1, 1))
	fmt.Println(defer3())
	fmt.Println(defer4())
}

func pointer() {
	var c int = 10
	var a *int = &c
	var b = 100

	fmt.Println(a, b)
	fmt.Println(*a)
	fmt.Println(&b)
}

func structAndString() {
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

func array() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(arr1)
	fmt.Println(arr1[5])
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr2))
}

// ByteSize is float64
type ByteSize float64

const (
	// 通过赋予空白标识符来忽略第一个值
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func test1() {
	fmt.Println(YB)
	var fs ByteSize = 100000
	fmt.Println(fs)
}
