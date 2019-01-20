package main

import (
	"fmt"
	"io/ioutil"
)

/*ReadMyData ...呵呵
*     
*/
func ReadMyData(){
	fmt.Println([]byte("ABCDEF"))
	data,err := ioutil.ReadFile("D:/workspace/dev/go/src/mygo/data.txt")
	if err!=nil{
		fmt.Printf("error: %s",err.Error())
	}
	fmt.Println(string(data))
}