package transform

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"
)

var spaceChar = byte(' ')
var transformConfigPath = string("./transform/transform_config.json")

func NewTransform() Transform {
	return Transform{}
}

type Transform struct {
}

func (t Transform) Process(input string, outputFile *os.File) error {


	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			// machine_serial_number, err := jsonparser.GetString(actVal, "_source", "machine_serial_number")
			// if err != nil {
			// 	fmt.Println("JSON parsing error: ", err)
			// 	return
			// }
			//fmt.Println("timestamp: " + machine_serial_number)
			var message=string("")
			// jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// 	name,err:= jsonparser.GetString(value, "[1]")
			// 	version,err:=jsonparser.GetString(value, "version")
			// 	fmt.Println(name,version)
			// 		}, "_source", "osquery_distributed_query_result","result")


					jsonparser.ObjectEach(actVal, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
        //fmt.Printf("Key: '%s ' Value: '%s '\n", string(key), string(value))
				message=message+string(key)+"="+string(value)+" "

				return nil
					}, "_source", "osquery_distributed_query_result","result","[0]")

					//fmt.Println(message)
			// calendarTime, err := jsonparser.GetString(actVal, "_source", "osquery_distributed_query_result","result","[0]")
			// if err != nil {
			// 	fmt.Println("JSON parsing error: ", err)
			// 	return
			// }
			// fmt.Println(calendarTime)
			 t.processMessage(message, outputFile)
			// if nil != err {
			// 	return
			// }

		}, "hits", "hits")

	return nil
}

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
