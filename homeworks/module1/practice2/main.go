package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			ch <- rand.Int()
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
