package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 1000

	//execute loop while i > 100
	for i > 100 {
		//get the random number between 1 & 1001
		i = rand.Intn(1000) + 1
		fmt.Println("i es", i)
		if i > 100 {
			fmt.Println("i es ", i, "sigue loopeando")
		}
	}
	fmt.Println("got", i, "se rompe el loop")
}
