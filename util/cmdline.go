package util

import (
	"flag"
)

func NewCmdLine() CmdLine {
	output := flag.String("output", "./"+file+".csv", "File to dump the transformed output.")
	c := new(CmdLine)
	c.outputPath = *output
	return *c
}

type CmdLine struct {
	outputPath string
}

func (c CmdLine) GetOutputPath() string {
	return c.outputPath
}
