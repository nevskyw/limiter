package balast

import "runtime"

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB memType = 1 << (10 * iota)
	MB
)

type (
	memType int64

	// Mem interface returns new
	Mem interface {
		Free()
	}

	balast struct {
		b *[]byte
	}
)

// New returns new Mem ballast controller
func New(count memType) Mem {
	b := make([]byte, count)
	return &balast{
		b: &b,
	}
}

func (b *balast) Free() {
	*b.b = nil
	runtime.GC()
}
