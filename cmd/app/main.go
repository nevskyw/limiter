package main

import (
	"limiter/internal/tools/balast"
	"limiter/internal/tools/graceful"
)

func main() {
	_ = graceful.Graceful()

	b := balast.New(balast.MB * 5)

	b.Free()
}
