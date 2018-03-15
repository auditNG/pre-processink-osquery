package transform

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"os"
	// "strconv"
	"strings"
)

const (
	GREEN = 1 + iota
	YELLOW
	RED
)

func NewFIMTransformer() FIMTransformer {
	return FIMTransformer{
		confObj: new(TransformConfig),
	}
}

type FIMTransformer struct {
	confObj *TransformConfig
}

func (f FIMTransformer) Init(config string) error {
	err := json.Unmarshal([]byte(config), f.confObj)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (f FIMTransformer) Process(input string, test string, config string, outputFile *os.File) error {
	err := f.Init(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	f.istestInWatchList(input,test,outputFile)
	return nil
}
func (f FIMTransformer) istestInWatchList(input string,test string,outputFile *os.File) {
	var configprobe=string("")
	for _,val := range f.confObj.Transform_config {
			configprobe=configprobe+val.Probe_name+"\n"
	}
	var esprobe=strings.Split(test,"\n")
	for _,val:=range esprobe{
		var check=strings.Contains(configprobe,val)
		if  check==true{
			f.applyRules(input,val,outputFile)
		} else {
			f.parseAndWrite(input,outputFile,val)
		}
	}
}
func (f FIMTransformer) parseAndWrite(input string,outputFile *os.File,test string) error {
	var message=string("")
	var check []byte
	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			check,_,_,err=jsonparser.Get(actVal,"_source", "osquery_distributed_query_result","probe","name")
			if string(check)==test{
				jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int,err error)  {
					jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
						message=message+string(pair)+"\n"
						return nil
					})
				var outputLine=strings.Join(strings.Split(message,"\n"), ",") + "\n"
				if _,err:=outputFile.Write([]byte(outputLine));err!=nil{
				 fmt.Println("Error weiting line to output file")
		 	 	}
				message=""
				}, "_source", "osquery_distributed_query_result","result")
		 }
	}, "hits","hits")
		return nil
}
<<<<<<< HEAD
func (f FIMTransformer) applyRules(input string,test string,outputFile *os.File){
	var message=string("")
=======
func (f FIMTransformer) applyRules(input string, test string) {
	var message = string("")
>>>>>>> 01e7799e48c1e9feb1afe40bd4c190f86faa3991
	var check []byte
	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			check, _, _, err = jsonparser.Get(actVal, "_source", "osquery_distributed_query_result", "probe", "name")
			if string(check) == test {
				jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					jsonparser.ObjectEach(value, func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error {
						message = message + string(key) + "=" + string(pair) + "\n"

						return nil
					})
					var outputLine=strings.Join(strings.Split(message,"\n"), ",") + "\n"
					if _,err:=outputFile.Write([]byte(outputLine));err!=nil{
					 fmt.Println("Error weiting line to output file")
				 }
				 message=""
						}, "_source", "osquery_distributed_query_result","result")

			}
			}, "hits","hits")

				}, "_source", "osquery_distributed_query_result", "result")

			}
		}, "hits", "hits")
}
