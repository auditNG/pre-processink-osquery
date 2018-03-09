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


	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			var message=string("")
			var data map[string]string
			data=make(map[string]string)
			// jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// 	name,err:= jsonparser.GetString(value, "[1]")
			// 	version,err:=jsonparser.GetString(value, "version")
			// 	fmt.Println(name,version)
			// 		}, "_source", "osquery_distributed_query_result","result")


			jsonparser.ObjectEach(actVal, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
      	//fmt.Printf("Key: '%s ' Value: '%s '\n", string(key), string(value))
				message=message+string(key)+"="+string(value)+"\n"
				data[string(key)]=string(value)

				return nil
					}, "_source", "osquery_distributed_query_result","result","[0]")
			//fmt.Println(message)
			for key,value:= range data{
				fmt.Println("key:",key,"  value:",value)
			}
			//t.extractMessage(message,outputFile)
			t.processMessage(message, outputFile)

		}, "hits", "hits")

	return nil
}
//func (t Transform) extractMessage(message string,outputFile *os.File) ([]string,error){
	//var content []string =strings.Fields(message)
	//var content []string =strings.Split(message,"\n")

	//return content,nil
//}

func (t Transform) processMessage(message string, outputFile *os.File) error {

	fimTransformer := NewFIMTransformer()

	config, err := ioutil.ReadFile(transformConfigPath)
	if err != nil {
		fmt.Println("Error reading transform config")
		return err
	}

	fimTransformer.Process(message, string(config), outputFile)

	return nil
}
