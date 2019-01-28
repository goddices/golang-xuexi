package language
import "fmt"

// ReturnTest is to try golang return
func ReturnTest(){
	voidFunc(true)
	voidFunc(false)
	fmt.Println(add(1,2))
	fmt.Println(operation(1,2))
	fmt.Println(operation(0,0))
	fmt.Println(operation2())
}

// return void
func voidFunc(a bool) {
	if a {
		fmt.Println("param a is true")
		return
	}
	fmt.Println("other ops")
}

// return the value of a+b
func add(a int,b int) int { 
	return a+b
}


func operation(a int,b int) (s int) {
	if a+b == 0{
		s = 0
		return // return s
	}
	return a+b // return a+b, 
	//也可以这样理解
	//s=a+b
	//return 
}

func operation2() (s int) {
	s++
	return 0 // return a+b, 
	//也可以这样理解
	//s=a+b
	//return 
}