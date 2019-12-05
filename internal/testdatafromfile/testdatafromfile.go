package testdatafromfile

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func From(filename string) *os.File {
	_, file, _, _ := runtime.Caller(0)
	f, err := os.Open(filepath.Join(path.Dir(file), "../../testdata", filename))
	if err != nil {
		panic(err)
	}
	return f
}
