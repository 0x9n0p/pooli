# pooli
Worker pool with dynamic size

```go
	p := pooli.Open(5)
	p.Start()

	p.SendTask(pooli.NewTask(func(ctx context.Context) error {
		fmt.Println("task")
		return nil
	}))

	p.SetGoroutines(1)
```
