package transform

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"strings"
)

var spaceChar = byte('\n')
var transformConfigPath = string("./transform/transform_config.json")

func NewTransform() Transform {
	return Transform{}
}

type Transform struct {
}

func(t Transform) ProcessMeta(input string) error{
	jsonparser.ArrayEach([]byte(input),func(actVal []byte, _ jsonparser.ValueType, _ int,err error){
		created_at,_:=jsonparser.GetString(actVal, "_source", "created_at")
		ip,_:=jsonparser.GetString(actVal, "_source", "request","ip")
		msn,_:=jsonparser.GetString(actVal,"_source","machine_serial_number")
		jsonparser.ArrayEach(actVal, func(value []byte, _ jsonparser.ValueType, _ int,err error)  {
		jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
		fmt.Println(string(pair))
		return nil
		})
	},"_source","machine","meta_business_units")
		fmt.Println(created_at,ip,msn)
	},"hits","hits")
	return nil
}
func (t Transform) Process(input string) error {
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
	t.processMessage(input,name)
	return nil
}
func (t Transform) processMessage(input string, test string) error {
	fimTransformer := NewFIMTransformer()
	config, err := ioutil.ReadFile(transformConfigPath)
	if err != nil {
		fmt.Println("Error reading transform config")
		return err
	}
	fimTransformer.Process(input,test, string(config))
	return nil
}
