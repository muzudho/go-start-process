# go-start-process

Go言語の練習用（＾～＾）  

外部プロセスを起動する（＾～＾）  

## Set up

```console
go build
```

## Run

```shell
go-start-process -FilePath <ExecutableFilePath> -ArgumentList <CommandLineParameters>
# Example:
#
# ```shell
# go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe
# ```
```

また、カレント・ディレクトリを移動したときに ファイルがどこに出力されるかも確認すること。  
ただし、以下の　実行ファイルを相対パスで指定する方法は、健全な拡張性を塞いでしまいます。  

```shell
# Example:
#
# ```shell
# powershell
# Set-Location C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace
# ```

../go-start-process -FilePath <ExecutableFilePath> -ArgumentList <CommandLineParameters>
# Example
# ../go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe
```
