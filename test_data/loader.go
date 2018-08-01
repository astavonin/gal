package test_data

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

func FromFile(name string) string {
	_, cur_path, _, _ := runtime.Caller(0)
	dir, _ := path.Split(cur_path)

	fname := fmt.Sprintf("%s/expected/%s.txt", dir, name)
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(data)
}
