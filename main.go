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
	fmt.Printf("DefaultWorkingDirectory=[%s]\n", dwd)

	// コマンドライン引数
	filePath := flag.String("FilePath", dwd, "Executable file path.")
	argumentList := flag.String("ArgumentList", "", "Parameters.")
	flag.Parse()
	fmt.Printf("flag.Args()=%s\n", flag.Args())

	parameters := strings.Split(*argumentList, " ")

	startProcess(*filePath, parameters)
}

func startProcess(filePath string, parameters []string) {
	cmd := exec.Command(filePath, parameters...)
	err := cmd.Start()
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}
	cmd.Wait()
}
