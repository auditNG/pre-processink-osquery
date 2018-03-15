package transform

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"
	"strings"

)

var spaceChar = byte('\n')
var transformConfigPath = string("./transform/transform_config.json")

func NewTransform() Transform {
	return Transform{}
}

type Transform struct {
}

func (t Transform) Process(input string, outputFile *os.File) error {
var test []byte
var name=string("")
jsonparser.ArrayEach([]byte(input),
	func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
		 test,_,_,err=jsonparser.Get(actVal,"_source", "osquery_distributed_query_result","probe","name")
		 var contain=strings.Contains(name,string(test))
		 if contain==false {
				name=name+string(test)+"\n"
			}
		}, "hits","hits")
	t.processMessage(input,name,outputFile)
	return nil
}

func (t Transform) processMessage(input string,test string, outputFile *os.File) error {
	fimTransformer := NewFIMTransformer()
	config, err := ioutil.ReadFile(transformConfigPath)
	if err != nil {
		fmt.Println("Error reading transform config")
		return err
	}
	fimTransformer.Process(input,test, string(config), outputFile)
	return nil
}
