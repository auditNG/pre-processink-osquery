package transform

import (
  "os"
)

type Transformer interface {
	Process(message string, config string, outputFile *os.File) error
}
