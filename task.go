package pooli

import "context"

type Task struct {
	execute func(ctx context.Context) error
	success func()
	fail    func(error)
	final   func()
}

func NewTask(execute func(ctx context.Context) error) Task {
	return Task{execute: execute}
}

func (t Task) Success(s func()) Task {
	t.success = s
	return t
}

func (t Task) Fail(f func(err error)) Task {
	t.fail = f
	return t
}

func (t Task) Final(f func()) Task {
	t.final = f
	return t
}
