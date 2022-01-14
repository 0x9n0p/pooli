package pooli

type Status uint

const (
	Idle     = Status(iota)
	Progress = Status(iota)
)
