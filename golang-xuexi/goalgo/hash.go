package goalgo

import (
	"fmt"
)

func HashTest() {
	fmt.Println(hashAdd1(100))
	fmt.Println(hashAdd1(101))
	fmt.Println(hashAdd1(102))

	fmt.Println(hashFNVHash1("hehehehe"))
	fmt.Println(hashFNVHash1("heheheha"))
	fmt.Println(hashFNVHash1("我操厉害了"))
}

/*
1. 加法Hash；
2. 位运算Hash；
3. 乘法Hash；
4. 除法Hash；
5. 查表Hash；
6. 混合Hash；
*/

//
func hashAdd1(input int) int {
	return (input % 8) ^ 5
}

func hashFNVHash1(data string) uint {
	var p uint = 16777619
	hash := uint(2166136261)
	for i := 0; i < len(data); i++ {
		hash = (hash ^ uint(data[i])) * p
	}
	hash += hash << 13
	hash ^= hash >> 7
	hash += hash << 3
	hash ^= hash >> 17
	hash += hash << 5
	return hash
}
