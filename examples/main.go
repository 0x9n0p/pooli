package main

import (
	"context"
	"fmt"
	"time"

	"github.com/0x9n0p/pooli"
)

func main() {
	p := pooli.Open(context.Background(), pooli.Config{
		Goroutines: 5,
		Pipe:       make(chan pooli.Task),
	})
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
