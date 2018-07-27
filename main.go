package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"

	"github.com/astavonin/gal/algo"
)

func main() {

	callTests()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test_data/test_file.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	//ast.Print(fset, node)

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err) // type error
	}
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {

			//a, ok := obj.Type().Underlying().(*types.Array)
			//if ok {
			//fmt.Println("Array: ", a.Elem())
			//}
			s, ok := obj.Type().Underlying().(*types.Slice)
			if ok {
				fmt.Print("Slice:", obj.Type(), "-->", s.Elem())
				fmt.Println(", cmp: ", types.Comparable(s.Elem()))

				g, _ := algo.NewGenerator(obj.Type())
				g.Generate()
			}
		}
	}

	//for _, decl := range f.Decls {
	//	gen, ok := decl.(*ast.GenDecl)
	//	if !ok {
	//		continue
	//	}
	//		switch spec := spec.(type) {
	//		case *ast.TypeSpec:
	//			switch t := spec.Type.(type) {
	//			case *ast.ArrayType:
	//				fmt.Println("ArrayType", spec.Name, t.Elt, reflect.TypeOf(t.Len))
	//			case *ast.MapType:
	//				fmt.Println("MapType", spec.Name, t.Key, t.Value)
	//			}
	//		}
	//	}
	//}

}

func callTests() {
	//s1 := test_data.TestStringSlice{"a", "ss", "b"}
	//fmt.Println("TestStringSlice:", s1.Find("ss"))
	//
	//s2 := test_data.TestStructSlice{
	//	test_data.TestStruct{1, "foo"},
	//	test_data.TestStruct{2, "boo"},
	//}
	//s2test := test_data.TestStruct{1, "foo"}
	//fmt.Println("TestStructSlice:", s2.Find(&s2test))
}
