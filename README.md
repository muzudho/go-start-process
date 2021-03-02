# go-start-process

Go言語の練習用（＾～＾）  

外部プロセスを起動する（＾～＾）  

## Set up

```console
go build
```

## Run

### Step 1

```shell
go-start-process -FilePath <ExecutableFilePath> -ArgumentList <CommandLineParameters>
# Example:
#
# ```shell
# go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe
# ```
```

### Step 2

また、カレント・ディレクトリを移動したときに ファイルがどこに出力されるかも確認してください。  
ただし、この実行ファイルを相対パスで指定する方法は、健全な拡張性を塞いでしまいます。  

```shell
# `powershell` と打鍵して、Posershell を使って見ましょう
powershell

Set-Location <WorkingDirectoryPath>
# Example
# Set-Location C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace-1

../go-start-process -FilePath <ExecutableFilePath> -ArgumentList <CommandLineParameters>
# ../go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe

# powershell から抜けてください
exit

# また、カレントディレクトリーを元に戻しておいてください
# cd ..
```

### Step 3

作業ディレクトリは指定できるべきですが、プロセス間の処理方法はＯＳごとに異なり一般的な方法がありません。  
呼び出される側のアプリケーションが、作業ディレクトリを受け取れるようにしておいてください。  

```shell
go-start-process -FilePath <ExecutableFilePath> -ArgumentList <CommandLineParameters>
# Example (Case of Windows)
# `^` - コマンドラインを押し返したいときは末尾に付けてください
# `"` - 引数が半角空白を含むときはダブルクォーテーションで囲んでください
# go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe ^
# -ArgumentList "-WorkingDirectory C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace-1"
```

### Step 4

では、 `go-start-process.exe` から `go-start-process.exe` を呼び出すとどうなるでしょうか？  
入れ子にして試してみましょう。  
パラメーターがどこで分割されて　どのような塊となって、ちゃんと渡されているかが論旨です。  

```shell
# 3行目が長いことに注意してください
go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/go-start-process.exe ^
-WorkingDirectory C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace-1 ^
-ArgumentList "-ArgumentList go-start-process -FilePath C:/Users/むずでょ/go/src/github.com/muzudho/go-count-up/go-count-up.exe -WorkingDirectory C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace-2 -ArgumentList \"-WorkingDirectory C:/Users/むずでょ/go/src/github.com/muzudho/go-start-process/workspace-3\""
```
