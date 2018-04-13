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

func (f FIMTransformer) Process(input string, test string, config string) error {
	err := f.Init(config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	f.istestInWatchList(input,test)
	return nil
}
func (f FIMTransformer) istestInWatchList(input string,test string) {
	var configprobe=string("")
	for _,val := range f.confObj.Transform_config {
			configprobe=configprobe+val.Probe_name+"\n"
	}
	var esprobe=strings.Split(test,"\n")
	for _,val:=range esprobe{
		var check=strings.Contains(configprobe,val)
		if  check==true{
			f.applyRules(input,val)
		} else {
			f.parseAndWrite(input,val)
		}
	}
}
func (f FIMTransformer) parseAndWrite(input string,test string) error {
	var message=string("")
	var fields=string("")
	var fcheck=true
	var check []byte
	d :=NewDataMapper()
	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			check,_,_,err=jsonparser.Get(actVal,"_source", "osquery_distributed_query_result","probe","name")
			created_at,_:=jsonparser.GetString(actVal, "_source", "created_at")
			ip,_:=jsonparser.GetString(actVal, "_source", "request","ip")
			msn,_:=jsonparser.GetString(actVal,"_source","machine_serial_number")
			if string(check)==test{
				f, err :=os.OpenFile(test+".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err!=nil{
					fmt.Println("Error: ", err)
        return
				}
				message=message+created_at+"\n"+ip+"\n"+msn+"\n"
				fields=fields+"created_at"+"\n"+"ip"+"\n"+"msn"+"\n"
				//parsing meta_business_units
				jsonparser.ArrayEach(actVal, func(value []byte, _ jsonparser.ValueType, _ int,err error)  {
				jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
				fields=fields+string(key)+"\n"
				message=message+string(pair)+"\n"
				return nil
				})
				},"_source","machine","meta_business_units")
				meta:=message
				//parsing actual data
				jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int,err error)  {
					jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
						message=message+string(pair)+"\n"
						fields=fields+string(key)+"\n"
						return nil
					})
				if fcheck==true{
					var fieldName=strings.Join(strings.Split(fields,"\n"),",")
					fieldName=strings.TrimRight(fieldName,",")
					fieldName=fieldName+"\n"
					if _,err:=f.Write([]byte(fieldName));err!=nil{
					 fmt.Println("Error writing line to output file")
			 	 	}
					fcheck=false

				}
				var outputLine=strings.Join(strings.Split(message,"\n"), ",")
				outputLine=strings.TrimRight(outputLine,",")
				outputLine=outputLine+"\n"
				d.mapper(message,fields)
				if _,err:=f.Write([]byte(outputLine));err!=nil{
				 fmt.Println("Error writing line to output file")
		 	 	}
				message=meta
				fields=""
				}, "_source", "osquery_distributed_query_result","result")
				message=""
		 }
	}, "hits","hits")

		return nil
}
func (f FIMTransformer) applyRules(input string,test string){
	var message=string("")
	var fields=string("")
	var fcheck=true
	var check []byte
	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			check, _, _, err = jsonparser.Get(actVal, "_source", "osquery_distributed_query_result", "probe", "name")
			created_at,_:=jsonparser.GetString(actVal, "_source", "created_at")
			ip,_:=jsonparser.GetString(actVal, "_source", "request","ip")
			msn,_:=jsonparser.GetString(actVal,"_source","machine_serial_number")
			if string(check) == test {
				f, err :=os.OpenFile(test+".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				if err!=nil{
					fmt.Println("Error: ", err)
        return
				}
				message=message+created_at+"\n"+ip+"\n"+msn+"\n"
				fields=fields+"created_at"+"\n"+"ip"+"\n"+"msn"+"\n"
				//parsing meta_business_units
				jsonparser.ArrayEach(actVal, func(value []byte, _ jsonparser.ValueType, _ int,err error)  {
				jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
				fields=fields+string(key)+"\n"
				message=message+string(pair)+"\n"
				return nil
				})
				},"_source","machine","meta_business_units")
				meta:=message
				//parsing actual data
				jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					jsonparser.ObjectEach(value, func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error {
						message = message + string(pair) + "\n"
						fields=fields+string(key)+"\n"

						return nil
					})
					if fcheck==true{
						var fieldName=strings.Join(strings.Split(fields,"\n"),",")+"\n"
						if _,err:=f.Write([]byte(fieldName));err!=nil{
						 fmt.Println("Error writing line to output file")
				 	 	}
						fcheck=false
					}
					var outputLine=strings.Join(strings.Split(message,"\n"), ",") + "\n"
					// d :=NewDataMapper()
					// d.mapper(message,fields)
					if _,err:=f.Write([]byte(outputLine));err!=nil{
					 fmt.Println("Error writing line to output file")
				 }
				 message=meta
				 fields=""
						}, "_source", "osquery_distributed_query_result","result")

			}
			}, "hits","hits")

}
