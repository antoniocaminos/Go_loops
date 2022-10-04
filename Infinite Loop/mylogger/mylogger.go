package mylogger

import "log"

func ListenFoLog(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}
