package transform

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  // "github.com/buger/jsonparser"
  "encoding/json"
)

func NewFIMTransformer() FIMTransformer {
  return FIMTransformer{}
}

type FIMTransformer struct{}

func (f FIMTransformer) Process(message string, config string, outputFile *os.File) error {

  transformHelper := NewTransformHelper()
  val, _ := f.getSyscalls(config)
  fmt.Println(val)


  syscall, err := transformHelper.GetIntValue(message, "syscall=")
	if nil != err {
		fmt.Println("Unable to get syscall")
	}
	fmt.Println("syscall: " + strconv.Itoa(syscall))

	exitcode, err := transformHelper.GetIntValue(message, "exit=")
	if nil != err {
		fmt.Println("Unable to get exitcode")
	}
	fmt.Println("exitcode: " + strconv.Itoa(exitcode))


  return nil
}

func (f FIMTransformer) getSyscalls(config string) ([]int, error) {
  var confObj TransformConfig
  err := json.Unmarshal([]byte(config), &confObj)

    if err != nil {
          fmt.Println("Unmarshall error")
          fmt.Println(err)
          return nil,nil
    }
    fmt.Printf("%+v\n", confObj)
    fmt.Println(confObj.Fim.SyscallList);
    fmt.Println(strings.Join(confObj.Fim.SyscallList, ", "))

    marsh, _ := json.Marshal(confObj)
    fmt.Println("Marshalled: " + string(marsh))
    return nil, nil

}
