package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	//read input from user 5 times and write a log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go mylogger.ListenFoLog(ch)

	fmt.Println("Enter Something")
	for i := 0; i < 5; i++ {
		fmt.Print("==> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
