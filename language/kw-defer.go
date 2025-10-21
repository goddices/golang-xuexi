package language

import "fmt"

// DeferTest is to try defer keyword
func DeferTest() {

	fmt.Println("first call in DeferTest")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("\nDeferTest about to return")
	fmt.Println("\nDeferTest deferFunction result:", deferFunction())

	fmt.Println("DeferTest about to return 2")
	fmt.Println("\nDeferTest deferFunction2 result:", deferFunction2())
}

func deferFunction() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func deferFunction2() int {
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
