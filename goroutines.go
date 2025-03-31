package main

import (
	"fmt"
	"time"
)

func greeting(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(name)
	}
}

func main() {
	go greeting("foo")
	greeting("bar")
}
