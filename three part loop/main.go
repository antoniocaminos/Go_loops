package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("i es ", i)
	}
	for i := 10; i >= 0; i-- {
		fmt.Println("i es ahora", i)
	}
}
