package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		for{
			a := 1
			fmt.Print(a)
			time.Sleep(1);
		}

	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int){
			fmt.Println("Hello world",i)
			wg.Done()
		}(i)
	}
	//wg.Wait()
}