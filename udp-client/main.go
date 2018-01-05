package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	addr := os.Args[1]
	log.Println("Connecting to", addr)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	eventTitle := "Some event"
	eventText := "Some text"
	tag := "app:upd-client"
	event := fmt.Sprintf("_e{%d,%d}:%s|%s|#%s", len(eventTitle), len(eventText), eventTitle, eventText, tag)
	for {
		message := []byte(event)
		count, err := conn.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		if count != len(message) {
			log.Println("Failed to send full message")
		}
		time.Sleep(time.Second)
	}
}
