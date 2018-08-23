package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

func main() {

	go pinger(irc)
	go sendMsg(sms, "Hi")
	go ponger(irc)
	go sendMsg(sms, "Hello")

	go printer()

	var exit string
	fmt.Scanln(&exit)
}

func pinger(channel chan string) {
	for {
		channel <- "ping"
	}
}

func ponger(channel chan string) {
	for {
		channel <- "pong"
	}
}

func sendMsg(channel chan string, msg string) {
	for {
		channel <- msg
	}
}

func printer() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Printf("irc:%s\n", msg)
		case msg = <-sms:
			fmt.Printf("sms:%s\n", msg)
		}
		time.Sleep(time.Second * 1)
	}
}
