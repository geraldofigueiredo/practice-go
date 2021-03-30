package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("Boring")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-c:
			fmt.Printf("You say: %q\n", <-c)
		case <-timeout:
			fmt.Println("timeout")
			os.Exit(1)
		}
	}
}
