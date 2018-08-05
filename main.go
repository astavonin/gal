package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"path/filepath"

	"go/build"

	"go/ast"

	"go/importer"

	"github.com/astavonin/gal/algo"
)

type StringSet map[string]struct{}

type Generator struct {
	dir   string
	defs  []string
	files []string
}

func newGenerator(dir string, defs []string) *Generator {
	return &Generator{
		dir:  dir,
		defs: defs,
	}
}

func toStringSet(strings []string) StringSet {
	set := make(StringSet)
	for _, name := range strings {
		set[name] = struct{}{}
	}
	return set
}

func (g *Generator) generate(typesList []string) error {

	err := g.buildFilesList()
	if err != nil {
		return err
	}
	fmt.Println(g.files)

	pkg, err := g.parse()
	if err != nil {
		return err
	}https://app.grammarly.com/
	typesSet := toStringSet(typesList)

	for _, name := range pkg.Scope().Names() {
		_, ok := typesSet[name]
		if !ok {
			continue
		}
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			switch t := obj.Type().Underlying().(type) {
			case *types.Slice:
				fmt.Print("Slice:", obj.Type(), "-->", t.Elem())
				fmt.Println(", cmp: ", types.Comparable(t.Elem()))

				sliceGen, _ := algo.NewGenerator(obj.Type())
				buf, err := sliceGen.Generate()
				if err != nil {
					return err
				}
				fmt.Println(buf)
			case *types.Array:
				fmt.Print("Array:", obj.Type(), "-->", t.Elem())
				fmt.Println(", cmp: ", types.Comparable(t.Elem()))
			}
		}
	}
	return nil
}

func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

func (g *Generator) buildFilesList() error {
	ctx := build.Default
	ctx.BuildTags = g.defs

	pkg, err := ctx.ImportDir(g.dir, 0)
	if err != nil {
		return fmt.Errorf("cannot process directory %s: %s", g.dir, err)
	}
	var files []string
	files = append(files, pkg.GoFiles...)
	files = append(files, pkg.CgoFiles...)
	files = append(files, pkg.SFiles...)

	g.files = prefixDirectory(g.dir, files)

	return nil
}

func (g *Generator) parse() (*types.Package, error) {

	fset := token.NewFileSet()
	var astFiles []*ast.File

	for _, fname := range g.files {
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			log.Fatalf("parsing error: %s, %s", fname, err)
		}
		astFiles = append(astFiles, f)
	}
	if len(astFiles) == 0 {
		return nil, fmt.Errorf("%s: ho Go files found", g.dir)
	}

	config := types.Config{
		IgnoreFuncBodies: true,
		Importer:         importer.For("source", nil),
		FakeImportC:      true,
	}

	return config.Check("", fset, astFiles, nil)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	var tags []string
	g := newGenerator(args[0], tags)
	g.generate([]string{"TestStructSlice", "TestStringSlice"})
}
