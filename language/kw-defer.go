package language
import "fmt"
// DeferTest is to try defer keyword
func DeferTest() { 
	defer1()
	fmt.Println()
	fmt.Println(defer2(1, 1))
	fmt.Println(defer3())
	fmt.Println(defer4())
}



func defer1() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	} 
	/* 被defer的函数，实际声明了5次，以栈的顺序执行（先进后出）：
	 * fmt.Printf("%d ", 4)  第五次defer
	 * fmt.Printf("%d ", 3)  第四次defer
	 * fmt.Printf("%d ", 2)  第三次defer
	 * fmt.Printf("%d ", 1)  第二次defer
	 * fmt.Printf("%d ", 0)  第一次defer
	 */
}

 
func defer2(x, y int) (z int) {
	//Go语言的return关键字,参见kw-return.go
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
		fmt.Println("deferB:", i) 
	}()
	defer func() {
		i++
		fmt.Println("deferA:", i) 
	}()
	return i
}
