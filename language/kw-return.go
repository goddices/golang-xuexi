package language
import "fmt"

// return void
func voidFunc(a bool) {
	if a {
		fmt.Println("a")
	}
	return
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