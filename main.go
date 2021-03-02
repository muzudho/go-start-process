package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
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
	wd := flag.String("WorkingDirectory", dwd, "Working directory path.")
	fmt.Printf("GSP     | WorkingDirectory=%s\n", *wd)

	filePath := flag.String("FilePath", dwd, "Executable file path.")
	fmt.Printf("GSP     | ExeFilePath=%s\n", *filePath)

	argumentList := flag.String("ArgumentList", "", "Parameters.")
	flag.Parse()

	parameters := strings.Split(*argumentList, " ")
	fmt.Printf("GSP     | Parameters=[%s]\n", strings.Join(parameters, " "))

	externalProcessLogName := filepath.Join(*wd, "external-process.log")
	fmt.Printf("GSP     | ExternalProcessLogName=%s\n", externalProcessLogName)

	fmt.Printf("GSP     | Start\n")
	startProcess(*filePath, parameters, externalProcessLogName)
	fmt.Printf("GSP     | End\n")
}

func startProcess(exeFilePath string, parameters []string, externalProcessLogName string) {

	cmd := exec.Command(exeFilePath, parameters...)

	externalProcessStdout, _ := cmd.StdoutPipe()
	defer externalProcessStdout.Close()

	err := cmd.Start()
	if err != nil {
		panic(fmt.Errorf("cmd.Start() --> [%s]", err))
	}

	receiveExternalProcessStdout(externalProcessStdout, externalProcessLogName)

	cmd.Wait()
}

func receiveExternalProcessStdout(externalProcessStdout io.ReadCloser, externalProcessLogName string) {
	var buffer [1]byte // これが満たされるまで待つ。1バイト。

	p := buffer[:]

	for {
		n, err := externalProcessStdout.Read(p)

		if nil != err {
			if fmt.Sprintf("%s", err) == "EOF" {
				// End of file
				return
			}

			panic(err)
		}

		if 0 < n {
			bytes := p[:n]

			// 思考エンジンから１文字送られてくるたび、表示。
			writeString(externalProcessLogName, string(bytes))
		}
	}
}

func writeString(fileName string, contents string) {
	// 上書き用ファイル
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// ログのファイル・オープン失敗
		panic(err)
	}
	defer file.Close()

	// 数字を書込
	file.WriteString(contents)
}
