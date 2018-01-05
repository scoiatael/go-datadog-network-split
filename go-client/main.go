package main

import (
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	addr := os.Args[1]
	log.Println("Connecting to", addr)
	c, err := statsd.New(addr)
	if err != nil {
		log.Fatal(err)
	}
	c.Tags = append(c.Tags, "app:go-client")
	for {
		err := c.SimpleEvent("Some event", "Some text")
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
