# pooli
Worker pool with dynamic size

```go
	p := pooli.Open(context.Background(), pooli.Config{
		Goroutines: 5,
		Pipe:       make(chan pooli.Task),
	})
	p.Start()

	p.SendTask(pooli.NewTask(func(ctx context.Context) error {
		fmt.Println("task")
		return nil
	}))

	p.SetGoroutines(1)
```
