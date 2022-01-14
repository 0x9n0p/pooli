package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tooghi/pooli"
)

func main() {
	p := pooli.Open(5)
	p.Start()

	for i := 0; i < 1000; i++ {
		go func(n int) {
			p.SendTask(pooli.NewTask(func(ctx context.Context) error {
				fmt.Println("task ", n)
				<-time.Tick(time.Second * 1)
				return nil
			}))
		}(i)
	}

	<-time.Tick(time.Second * 3)
	p.SetGoroutines(1)

	// p.Close()

	select {}
}
