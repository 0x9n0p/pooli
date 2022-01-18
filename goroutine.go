package pooli

import (
	"context"
	"sync"
)

type Goroutine struct {
	status        Status
	Pipe          chan Task

	ctx context.Context
	cnl context.CancelFunc

	m *sync.RWMutex
}

func NewGoroutine(ctx context.Context, pipe chan Task) *Goroutine {
	ctx, cnl := context.WithCancel(ctx)
	return &Goroutine{
		status: Idle,
		Pipe:   pipe,

		ctx: ctx,
		cnl: cnl,

		m: new(sync.RWMutex),
	}
}

func (g *Goroutine) Start() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		wg.Done()
		for {
			select {
			case <-g.ctx.Done():
				return
			case t := <-g.Pipe:
				g.SetStatus(Progress)
				ExecuteTask(g.ctx, t)
				g.SetStatus(Idle)
			}
		}
	}()

	wg.Wait()
}

func (g *Goroutine) SetStatus(status Status) {
	g.m.Lock()
	defer g.m.Unlock()

	g.status = status
}

func (g Goroutine) Status() Status {
	return g.status
}

func (g *Goroutine) Kill() {
	g.cnl()
}
