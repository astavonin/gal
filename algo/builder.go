package algo

import (
	"fmt"
	"go/types"
)

type Generator interface {
	Generate() (string, error)
}

func NewGenerator(t types.Type) (g Generator, err error) {
	s, ok := t.Underlying().(*types.Slice)
	if ok {
		fmt.Print("Slice:", t, "-->", s.Elem())
		fmt.Println(", cmp: ", types.Comparable(s.Elem()))

		g, err = NewSliceGenerator(t, s.Elem())
	}
	return
}
