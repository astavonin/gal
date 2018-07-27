package algo

import (
	"bytes"
	"go/types"

	j "github.com/dave/jennifer/jen"
)

type SliceGenerator struct {
	isBasic   bool // inputType is Go basic type
	isPointer bool // inputType is pointer

	self      *j.Statement
	inputType *j.Statement
	cont      *j.Statement
	valCmp    *j.Statement
}

func initTypes(isBasic bool, dataType string) (comp, inputType *j.Statement) {
	comp = j.Id("val")
	inputType = j.Id(dataType)
	if !isBasic {
		comp = j.Op("*").Add(comp)
		inputType = j.Op("*").Add(inputType)
	}
	return
}

func NewSliceGenerator(tc, tv types.Type) (*SliceGenerator, error) {
	ptr, isPointer := tv.(*types.Pointer)
	_, isBasic := tv.(*types.Basic)
	dataType := tv.String()
	if isPointer {
		dataType = ptr.String()
	}

	comp, inputType := initTypes(isBasic, dataType)
	g := SliceGenerator{
		isBasic:   isBasic,
		isPointer: isPointer,

		self:      j.Id("s"),
		inputType: inputType,
		cont:      j.Op("*").Id(tc.String()),
		valCmp:    comp,
	}

	return &g, nil
}

func (s *SliceGenerator) Generate() (res string, err error) {
	//fmt.Println("--->")
	f := j.NewFile("test")
	find := s.genFind()
	rfind := s.genRFind()

	f.Add(find)
	f.Add(rfind)

	buf := &bytes.Buffer{}
	err = f.Render(buf)
	if err != nil {
		return
	}
	res = buf.String()

	return
}

func (s *SliceGenerator) genFindCommon(name string, loop *j.Statement) *j.Statement {
	fn := j.Func().Params(
		s.self.Clone().Add(s.cont)).Id(name).Params(
		j.Id("val").Add(s.inputType)).Add(j.Int())

	elAcc := j.Parens(j.Op("*").Add(s.self)).Index(j.Id("i"))
	if s.isPointer {
		elAcc = j.Op("*").Add(elAcc)
	}

	body := j.Block(
		j.Id("pos").Op(":=").Lit(-1),
		loop.Block(
			j.If(
				elAcc.Op("==").Add(s.valCmp).Block(
					j.Id("pos").Op("=").Id("i"),
					j.Break(),
				),
			),
		),
		j.Return(j.Id("pos")),
	)

	return fn.Add(body)
}

func (s *SliceGenerator) genFind() *j.Statement {
	return s.genFindCommon("Find",
		j.For(
			j.Id("i").Op(":=").Lit(0),
			j.Id("i").Op("<").Len(j.Op("*").Add(s.self)),
			j.Id("i").Op("++"),
		))
}

func (s *SliceGenerator) genRFind() *j.Statement {
	return s.genFindCommon("RFind",
		j.For(
			j.Id("i").Op(":=").Len(j.Op("*").Add(s.self)),
			j.Id("i").Op(">=").Lit(0),
			j.Id("i").Op("--"),
		))
}
