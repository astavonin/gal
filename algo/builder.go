package algo

import (
	"fmt"
	"go/types"
)

type Generator interface {
	Generate() (string, error)
}

func NewGenerator(t types.Type, pkgName string) (g Generator, err error) {
	s, ok := t.Underlying().(*types.Slice)
	if ok {
		if !types.Comparable(s.Elem()) {
			return nil, fmt.Errorf("%s is not comparable type", s.Elem().String())
		}
		g, err = NewSliceGenerator(t, s.Elem(), pkgName)
	}
	return
}
