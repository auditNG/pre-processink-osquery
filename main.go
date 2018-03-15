package main

import (
	"fmt"
	"github.com/pre-processink/source"
	"github.com/pre-processink/transform"
	"github.com/pre-processink/util"
	"os"
)

func main() {
	// Get the output file path from command line argument
	cmdLine := util.NewCmdLine("output")
	outputPath := cmdLine.GetOutputPath()
	outputFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if nil != err {
		fmt.Println("Error opening output file for appending.")
		return
	}

	// Read from the source
	var s source.Source = source.NewESSource()
	result, err := s.Fetch()

	if err != nil {
		fmt.Println("Error fetching from source")
		return
	}

	// Transform and write to output
	var t transform.Transform = transform.NewTransform()
	err = t.Process(result, outputFile)

	if err != nil {
		fmt.Println("Error fetching from source")
		return
	}
}
