package language

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// GoTest is to try go
func GoTest() {
	fmt.Println("go go go!")
	signal := make(chan bool)

	go worker("hello", signal)

	<-signal
	fmt.Println("done")
}

// ChanTest is to try chan
func ChanTest() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println("ChanTest start")
	// fmt.Println(<-messages)
	fmt.Println(<-messages)

	fmt.Println("ChanTest done")
}

func worker(s string, signal chan<- bool) {
	fmt.Println(s)
	signal <- true
}

func work2(s string, signal chan<- bool) {
	fmt.Println(s)
	signal <- true
	// s := <-signal // error: signal is sent only channel
}

func work3(s string, signal <-chan bool) {
	fmt.Println(s)
	<-signal
	// signal <- true // error: signal is receive only channel
}

// SelectTest is to try select
func SelectTest() {
	messages := make(chan string, 16)
	go func() { messages <- "test select 11111" }()
	go func() { messages <- "test select 22222" }()

	for i := 0; i < 2; i++ {
		time.After(time.Second * 1)
		select {
		case msg1 := <-messages:
			fmt.Println("msg1:", msg1)
		case msg2 := <-messages:
			fmt.Println("msg2:", msg2)
		}
	}
}

func ChanBlockedTest() {
	messages := make(chan string)
	go func() {
		fmt.Println("Please input your message:")
		var message string
		reader := bufio.NewReader(os.Stdin)
		message, _ = reader.ReadString('\n')
		messages <- message
		fmt.Println("sent message")
	}()
	fmt.Print("waiting for message...\n")
	msg := <-messages
	fmt.Println("received message:", msg)
	close(messages)
}
