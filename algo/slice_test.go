package algo

import (
	"testing"

	"fmt"

	j "github.com/dave/jennifer/jen"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSliceGenerator_initTypes(t *testing.T) {

	Convey("comparator should have correct type", t, func() {
		Convey("pointer and basic type", func() {
			comp, inputType := initTypes(true, "string")

			So(fmt.Sprintf("%#v", comp), ShouldEqual, "val")
			So(fmt.Sprintf("%#v", inputType), ShouldEqual, "string")
		})
		Convey("non-pointer and basic type", func() {
			comp, inputType := initTypes(true, "string")

			So(fmt.Sprintf("%#v", comp), ShouldEqual, "val")
			So(fmt.Sprintf("%#v", inputType), ShouldEqual, "string")
		})
		Convey("pointer and non-basic type", func() {
			comp, inputType := initTypes(false, "TestStruct")

			So(fmt.Sprintf("%#v", comp), ShouldEqual, "*val")
			So(fmt.Sprintf("%#v", inputType), ShouldEqual, "*TestStruct")
		})
		Convey("non-pointer and non-basic type", func() {
			comp, inputType := initTypes(false, "TestStruct")

			So(fmt.Sprintf("%#v", comp), ShouldEqual, "*val")
			So(fmt.Sprintf("%#v", inputType), ShouldEqual, "*TestStruct")
		})
	})
}

func TestSliceGenerator_genFind(t *testing.T) {

	Convey("[]string case", t, func() {
		var (
			find = `func (s *TestStringSlice) Find(val string) int {
	pos := -1
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == val {
			pos = i
			break
		}
	}
	return pos
}`
			rfind = `func (s *TestStringSlice) RFind(val string) int {
	pos := -1
	for i := len(*s); i >= 0; i-- {
		if (*s)[i] == val {
			pos = i
			break
		}
	}
	return pos
}`
		)
		comp, inputType := initTypes(true, "string")

		g := SliceGenerator{
			isPointer: false,
			isBasic:   true,

			self:      j.Id("s"),
			inputType: inputType,
			cont:      j.Op("*").Id("TestStringSlice"),
			valCmp:    comp,
		}
		res := fmt.Sprintf("%#v", g.genFind())
		So(res, ShouldEqual, find)

		res = fmt.Sprintf("%#v", g.genRFind())
		So(res, ShouldEqual, rfind)
	})
	Convey("[]*string case", t, func() {
		var (
			find = `func (s *TestStringSlice) Find(val string) int {
	pos := -1
	for i := 0; i < len(*s); i++ {
		if *(*s)[i] == val {
			pos = i
			break
		}
	}
	return pos
}`
			rfind = `func (s *TestStringSlice) RFind(val string) int {
	pos := -1
	for i := len(*s); i >= 0; i-- {
		if *(*s)[i] == val {
			pos = i
			break
		}
	}
	return pos
}`
		)
		comp, inputType := initTypes(true, "string")

		g := SliceGenerator{
			isPointer: true,
			isBasic:   true,

			self:      j.Id("s"),
			inputType: inputType,
			cont:      j.Op("*").Id("TestStringSlice"),
			valCmp:    comp,
		}
		res := fmt.Sprintf("%#v", g.genFind())
		So(res, ShouldEqual, find)

		res = fmt.Sprintf("%#v", g.genRFind())
		So(res, ShouldEqual, rfind)
	})
	Convey("[]TestStruct case", t, func() {
		var (
			find = `func (s *TestStructSlice) Find(val *TestStruct) int {
	pos := -1
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == *val {
			pos = i
			break
		}
	}
	return pos
}`
			rfind = `func (s *TestStructSlice) RFind(val *TestStruct) int {
	pos := -1
	for i := len(*s); i >= 0; i-- {
		if (*s)[i] == *val {
			pos = i
			break
		}
	}
	return pos
}`
		)
		comp, inputType := initTypes(false, "TestStruct")

		g := SliceGenerator{
			isPointer: false,
			isBasic:   false,

			self:      j.Id("s"),
			inputType: inputType,
			cont:      j.Op("*").Id("TestStructSlice"),
			valCmp:    comp,
		}
		res := fmt.Sprintf("%#v", g.genFind())
		So(res, ShouldEqual, find)

		res = fmt.Sprintf("%#v", g.genRFind())
		So(res, ShouldEqual, rfind)
	})
	Convey("[]*TestStruct case", t, func() {
		var (
			find = `func (s *TestStructSlice) Find(val *TestStruct) int {
	pos := -1
	for i := 0; i < len(*s); i++ {
		if *(*s)[i] == *val {
			pos = i
			break
		}
	}
	return pos
}`
			rfind = `func (s *TestStructSlice) RFind(val *TestStruct) int {
	pos := -1
	for i := len(*s); i >= 0; i-- {
		if *(*s)[i] == *val {
			pos = i
			break
		}
	}
	return pos
}`
		)
		comp, inputType := initTypes(false, "TestStruct")

		g := SliceGenerator{
			isPointer: true,
			isBasic:   false,

			self:      j.Id("s"),
			inputType: inputType,
			cont:      j.Op("*").Id("TestStructSlice"),
			valCmp:    comp,
		}
		res := fmt.Sprintf("%#v", g.genFind())
		So(res, ShouldEqual, find)

		res = fmt.Sprintf("%#v", g.genRFind())
		So(res, ShouldEqual, rfind)
	})
}
