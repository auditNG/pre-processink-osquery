package transform

import (
  "fmt"
  "os"
  "encoding/json"
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

type FIMTransformer struct{
  confObj *TransformConfig
}

func (f FIMTransformer) Init(config string) error {
  err := json.Unmarshal([]byte(config), f.confObj)
  if err != nil {
        fmt.Println(err)
  }
  return err
}

func (f FIMTransformer) Process(message string, config string, outputFile *os.File) error {
  // Read in config
  err := f.Init(config)
  if(err != nil) {
    fmt.Println(err)
    return nil
  }

  // Instantiate transform helper
  transformHelper := NewTransformHelper()

  //Get syscall
  syscall, err := transformHelper.GetIntValue(message, "syscall=")
	if nil != err {
		fmt.Println("Unable to get syscall")
    return nil
	}
  fmt.Println(syscall)

  //Get exitcode
	exitcode, err := transformHelper.GetIntValue(message, "exit=")
	if nil != err {
		fmt.Println("Unable to get exitcode")
    return nil
	}
  fmt.Println(exitcode)

  //Get executable
  executable, err := transformHelper.GetIntValue(message, "exe=")
  if nil != err {
    fmt.Println("Unable to get exitcode")
    return nil
  }
  fmt.Println(executable)

  //Get user
  user, err := transformHelper.GetIntValue(message, "user=")
	if nil != err {
		fmt.Println("Unable to get exitcode")
    return nil
	}
  fmt.Println(user)


  return nil
}

func (f FIMTransformer) isSyscallsInList(syscall int) (bool) {
    return false
}

func (f FIMTransformer) applyLabelAlgo(syscall int, exitcode int, executable string, user int) {

}
