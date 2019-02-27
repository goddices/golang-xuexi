package language

import ("fmt" 
"time" )

 

// GoTest is to try go 
func GoTest(){
	fmt.Println("go go go!")
	signal := make(chan bool)
	
	go worker("hello",signal)
 
    <-signal
    fmt.Println("done")
}

// ChanTest is to try chan
func ChanTest(){
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func worker(s string,signal chan<- bool){
	fmt.Println(s)
	signal<-true
}

// SelectTest is to try select
func SelectTest(){
	messages:=make(chan string,16)
	go func(){messages<-"test select 11111"}()
	go func(){messages<-"test select 22222"}()
	 
	for i:=0;i<2;i++ {
		time.After(time.Second  * 1)
		select {
		case msg1 := <-messages:
			fmt.Println("msg1:",msg1)
		case msg2 := <-messages:
			fmt.Println("msg2:",msg2)
		}
	}
}
