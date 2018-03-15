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
	f.istestInWatchList(input, test)
	return nil
}
func (f FIMTransformer) istestInWatchList(input string, test string) {
	if strings.Contains(test, "kernel_info") {
		f.applyRules(input, "kernel_info")
	}
}
func (f FIMTransformer) applyRules(input string, test string) {
	var message = string("")
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

				}, "_source", "osquery_distributed_query_result", "result")

			}
		}, "hits", "hits")
	fmt.Println(message)

}
