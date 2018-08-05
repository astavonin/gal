package go_test_files

type TestStruct struct {
	A int
	B string
}

type TestStructSlice []TestStruct

//type TestStringSlice []string

//type TestStructPtrSlice []*TestStruct
//type TestStringPtrSlice []*string

//type TestFloatSlice [32]float64
//type TestFloatMap map[int]float64
//
//func (self *TestStringSlice) Find(val string) int {
//	pos := -1
//	for i := 0; i < len(*self); i++ {
//		if (*self)[i] == val {
//			pos = i
//			break
//		}
//	}
//	return pos
//}

//func (s *TestStringSlice) RFind(val string) int {
//	pos := -1
//	for i := len(*s); i >= 0; i-- {
//		if (*s)[i] == val {
//			pos = i
//			break
//		}
//	}
//	return pos
//}

//
//func (s *TestStructSlice) Find(val *TestStruct) int {
//	pos := -1
//	for i := 0; i < len(*s); i++ {
//		if (*s)[i] == *val {
//			pos = i
//			break
//		}
//	}
//	return pos
//}

//
//func (s *TestStructPtrSlice) Find(val *TestStruct) int {
//	pos := -1
//	for i := 0; i < len(*s); i++ {
//		if *(*s)[i] == *val {
//			pos = i
//			break
//		}
//	}
//	return pos
//}
//
//func (s *TestStringPtrSlice) Find(val string) int {
//	pos := -1
//	for i := 0; i < len(*s); i++ {
//		if *(*s)[i] == val {
//			pos = i
//			break
//		}
//	}
//	return pos
//}
