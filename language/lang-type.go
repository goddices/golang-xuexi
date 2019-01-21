package language

import "fmt"
 
 
// TypeTest is to try type system
func TypeTest(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Printf("value is string = %s\n", str)
	case int:
		fmt.Printf("value is int = %d\n", str)
	}
}

// PointerTest is to try pointer *int 
func PointerTest() {
	var c int = 10
	var a *int = &c
	var b = 100

	fmt.Println(a, b)
	fmt.Println(*a)
	fmt.Println(&b)
}
 
// ArrayTest is to try array
func ArrayTest() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(arr1)
	fmt.Println(arr1[5])
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr2))
}

// StructTest is to try struct, class... 
func StructTest() {
	fmt.Println(YB)
	var fs ByteSize = 100000
	fmt.Println(fs)
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

