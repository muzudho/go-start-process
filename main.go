package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Default working directory
	dwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GSP     | DefaultWorkingDirectory=%s\n", dwd)

	// コマンドライン引数
	filePath := flag.String("FilePath", dwd, "Executable file path.")
	argumentList := flag.String("ArgumentList", "", "Parameters.")
	flag.Parse()
	fmt.Printf("GSP     | flag.Args()=%s\n", flag.Args())

	parameters := strings.Split(*argumentList, " ")

	startProcess(*filePath, parameters)
}

func startProcess(filePath string, parameters []string) {
	fmt.Printf("GSP     | filePath=%s\n", filePath)
	fmt.Printf("GSP     | parameters=%s\n", parameters)
	cmd := exec.Command(filePath, parameters...)
	err := cmd.Start()
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}
	cmd.Wait()
}
