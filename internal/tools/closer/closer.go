package closer

import (
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
)

var Closer = New()
var instance *closer

type closer struct {
	closers *[]func() error
}

func New() *closer {
	o := sync.Once{}
	o.Do(func() {
		c := make([]func() error, 0)
		instance = &closer{
			closers: &c,
		}
	})
	return instance
}

func (c *closer) Add(f ...func() error) {
	*c.closers = append(*c.closers, f...)
}

func (c *closer) Close() {

	g := &errgroup.Group{}

	for _, f := range *c.closers {
		g.Go(f)
	}

	if err := g.Wait(); err != nil {
		log.Println(err)
	}

}
