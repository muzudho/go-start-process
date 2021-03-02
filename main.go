package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Working directory
	wdir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("wdir=[%s] err=[%s]", wdir, err))
	}
	fmt.Printf("wdir=[%s]\n", wdir)

	// コマンドライン引数
	filePath := flag.String("FilePath", wdir, "Executable file path.")
	argumentList := flag.String("ArgumentList", "", "Parameters.")
	flag.Parse()
	fmt.Printf("flag.Args()=%s\n", flag.Args())

	parameters := strings.Split(*argumentList, " ")
	cmd := exec.Command(*filePath, parameters...)
	err = cmd.Start()
	if err != nil {
		panic(fmt.Errorf("cmd.Run() --> [%s]", err))
	}
	cmd.Wait()
}
