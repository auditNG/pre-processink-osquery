package transform

import (
	"fmt"
	"os"
	"path/filepath"
  "strings"
)
func NewTest() Test {
	return Test{}
}

type Test struct {
}

func(t Test) run() ([]string, error) {
	searchDir := "../"

	fileList := make([]string, 0)
  files:=make([]string,0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
    if strings.Contains(path,".csv")!=false{
    fileList = append(fileList, path)
    }
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
    files=append(files,strings.TrimPrefix(file,"../"))
	}
  for _,val:=range files{
    fmt.Println(val)
  }

	return fileList, nil
}
