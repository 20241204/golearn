# golang 环境安装说明
go 程序下载安装: https://go.dev/dl/  
visualstudio 编辑工具(免费)下载安装: https://code.visualstudio.com/download  
goland idea 编辑工具(收费)下载安装: https://www.jetbrains.com/zh-cn/go/download/other.html  
jetbra 破解站点(破解软件可能侵权): https://3.jetbra.in/  

## Golang 环境配置
### 打开 powershell 
    # 查看 go 版本命令
        go version
    # 查看 go 环境命令
        go env
    # 查看 GOPATH 官方路径命令
    # 官方默认为 %USERPROFILE%\go 
    # 即 C:\Users\你的用户名\go 
    # 在 powershell 中表示为 $env:USERPROFILE\go
        go env | findstr /IR 'GOPATH'  
        ```回显信息
        set GOPATH=C:\Users\你的用户名\go
        ```

### 打开 powershell 创建 GOPATH 自定义工作路径，比如用 D:\GolandProjects 作为工作目录
    # 创建 D:\GolandProjects 工作目录命令
    # src 存放源代码
    # pkg 存放已编译的库文件
    # bin 存放已编译的二进制文件
        mkdir -Force -Verbose D:\GolandProjects\src
        mkdir -Force -Verbose D:\GolandProjects\pkg
        mkdir -Force -Verbose D:\GolandProjects\bin

### 打开 powershell 替换官方 GOPATH 环境变量为自定义工作路径(同理还原环境变量也是同样的方法)
    # 注意接下来的操作很危险，可能破坏系统本身命令调用，一定要确定理解自己在做什么
    # 不确定的可以放弃使用默认的，等自己理解了再回来，没必要折磨自己

    # 运行以下命令来获取当前的用户级别 GOPATH 环境变量
        $userGoPath = [Environment]::GetEnvironmentVariable("GOPATH", "User")
    # 如果要替换用户级别 GOPATH 环境变量中的路径，运行以下命令
        $userGoPath = $userGoPath.Replace("$env:USERPROFILE\go", "D:\GolandProjects")
    # 打印环境变量，检查是否替换了 $env:USERPROFILE\go
        echo $userGoPath
    # 更新用户级别 GOPATH 环境变量，运行以下命令
        [Environment]::SetEnvironmentVariable("GOPATH", $userGoPath, "User")

    # 运行以下命令来获取当前的用户级别 Path 环境变量
        $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    # 如果要替换用户级别 GOPATH 环境变量中的路径，运行以下命令
        $userPath = $userPath.Replace(";$env:USERPROFILE\go\bin", ";D:\GolandProjects\bin")
    # 打印环境变量，检查是否替换了 $env:USERPROFILE\go\bin
        echo $userPath
    # 更新用户级别 GOPATH 环境变量，运行以下命令
        [Environment]::SetEnvironmentVariable("Path", $userPath, "User")

    # 关闭重开 powershell 检查 GOPATH 环境变量工作路径，如果打印了自定义路径则代表成功
        go env | findstr /IR 'GOPATH'
        ```回显信息
        set GOPATH=D:\GolandProjects
        ```
    # 获取 PATH 环境变量的值
        $pathEnv = $env:Path

    # 将 Path 环境变量的值按分号分割为数组
        $pathArray = $pathEnv -split ';'

    # 要查找的路径 D:\GolandProjects\bin
        $targetPath = "D:\GolandProjects\bin"

    # 查找并输出结果
        if ($pathArray -contains $targetPath) {
            Write-Output "路径已经写入在用户级别 Path 环境变量中."
        } else {
            Write-Output "路径不在用户级别 Path 环境变量中写入失败."
        }

    # 如果需要提取匹配的路径用来验证
        $foundPaths = $pathArray | Where-Object { $_ -eq $targetPath }
        $foundPaths
        ```回显信息
        D:\GolandProjects\bin
        ```

### 个人开发目录结构介绍
    ```个人代码项目结构
    GolandProjects
    └── src
        └── golearn
            └── 01day
    ```
    ```总结
    Go工作目录
    └── 源码目录
        └── 项目名
            └── 模块名
    ```

### 企业开发目录结构介绍
    ```企业代码项目结构
    GolandProjects
    └── src
        └── www.github.com
            └── 20241204
                └── golearn
                    └── 01day
    ```
    ```总结
    Go工作目录
    └── 源码目录
        └── 公司域
            └── 开发者/机构
                └── 项目名
                    └── 模块名
    ```

### 打开 powershell 创建自己的企业代码项目目录
### 比如 D:\GolandProjects\src\www.github.com\20241204\golearn
### 可选，没必要使用命令创建，用文件管理器也可以
    # 创建项目 golearn 目录命令
        mkdir -Force -Verbose D:\GolandProjects\src\www.github.com\20241204\golearn
    # 创建项目模块 01day 目录命令
        mkdir -Force -Verbose D:\GolandProjects\src\www.github.com\20241204\golearn\01day

### 为项目初始化创建 go.mod 文件
    # 命令
        cd D:\GolandProjects\src\www.github.com\20241204\golearn
        go mod init

### go build 编译
    # 打开 powershell 进入 main.go 文件所在目录
    # 假设在 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
    # 默认生成二进制文件名字为 .go 源码文件所在文件夹的名字
    # 编译二进制文件命令
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
        go build
    # 也自定义编译二进制文件名字，比如 test.exe (可选)
    # 注意 windows 可执行文件后缀为 .exe 其它系统可执行文件后缀可不是 .exe
        go build -o test.exe
    # 执行命令之后就能看到 01var.exe 或 test.exe 二进制文件在当前目录中生成
    # 可以在 poweshell 中直接执行看效果
        .\01var.exe
        # 或者
        .\test.exe

### go run 直接运行源码
    # 打开 powershell 进入 main.go 文件所在目录
    # 假设在 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
    # 编译二进制文件命令(默认生成二进制文件名字为所在文件夹的名字)
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
        go run main.go

### go install -v 编译二进制文件并移动到工作目录的 bin 目录
    # 打开 powershell 进入 main.go 文件所在目录
    # 比如 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
    # 编译二进制文件命令(默认生成二进制文件名字为所在文件夹的名字)
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
        go install -v
    # 此时二进制文件不在当前目录了，却仍可以执行
        01var.exe

### 跨平台编译
    # 打开 powershell 进入 main.go 文件所在目录
    # 假设在 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
    # 编译二进制文件命令(默认生成二进制文件名字为所在文件夹的名字)
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
    # 比如为 linux x86_64/amd64 系统编译
        # 设置临时环境变量并执行 go build
        $env:CGO_ENABLED = "0"
        $env:GOOS = "linux"
        $env:GOARCH = "amd64"
        # 编译 Go 程序
        go build
        # 恢复环境变量
        $env:CGO_ENABLED = "0"
        $env:GOOS = "windows"
        $env:GOARCH = "amd64"

    # 在一行中设置环境变量并执行 go build (可选)
        $env:CGO_ENABLED = "0"; $env:GOOS = "linux"; $env:GOARCH = "amd64" ; go build ; $env:CGO_ENABLED = "0"; $env:GOOS = "windows"; $env:GOARCH = "amd64"
    
    # 执行命令之后就能看到 01var 二进制文件在当前目录中生成
    # 将 01var 二进制文件复制到 linux x86_64/amd64 系统中赋予可执行权限就可以执行看效果
        chmod -v +x 01var
        ./01var

### go fmt 格式化代码
    # 打开 powershell 进入 main.go 文件所在目录
    # 比如 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
        go fmt main.go
    # 此时 main.go 文件的代码就被格式化了

### go doc 查看内置函数用法
    # 打开 powershell 进入 main.go 文件所在目录
    # 比如 D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var\main.go
        cd D:\GolandProjects\src\www.github.com\20241204\golearn\01day\01var
        go doc builtin.delete
    # 此时可以查看 delete 内置函数的源码和用法说明

### 打开 powershell 编译安装 GO 常用工具
#### Go 工具介绍
1. **dlv (Delve)**
   - **地址**: `github.com/go-delve/delve`
   - **简介**: `dlv` 是 Go 语言的调试器，为开发者提供强大的调试功能。
   - **安装**:
     ```bash
     go install -v github.com/go-delve/delve/cmd/dlv@latest
     ```

2. **fillstruct**
   - **地址**: `github.com/davidrjenni/reftools/cmd/fillstruct`
   - **简介**: `fillstruct` 自动填充 Go 结构体的字段。
   - **安装**:
     ```bash
     go install -v github.com/davidrjenni/reftools/cmd/fillstruct@latest
     ```

3. **gocode**
   - **地址**: `github.com/mdempsky/gocode` 和 `github.com/stamblerre/gocode`
   - **简介**: `gocode` 是代码自动补全工具，有两个版本可供选择。
   - **安装**:
     ```bash
     go install -v github.com/mdempsky/gocode@latest
     go install -v github.com/stamblerre/gocode@latest
     ```

4. **godef**
   - **地址**: `github.com/rogpeppe/godef`
   - **简介**: `godef` 用于快速跳转到代码中的定义位置。
   - **安装**:
     ```bash
     go install -v github.com/rogpeppe/godef@latest
     ```

5. **goplay**
   - **地址**: `github.com/haya14busa/goplay`
   - **简介**: `goplay` 将 Go 代码片段发送到 Go Playground 上运行。
   - **安装**:
     ```bash
     go install -v github.com/haya14busa/goplay/cmd/goplay@latest
     ```

6. **goline**
   - **地址**: `golang.org/x/lint/golint`
   - **简介**: `golint` 提供 Go 代码风格检查建议。
   - **安装**:
     ```bash
     go install -v golang.org/x/lint/golint@latest
     ```

7. **gopkgs**
   - **地址**: `github.com/uudashr/gopkgs/v2/cmd/gopkgs`
   - **简介**: `gopkgs` 列出当前工作区或模块下的所有 Go 包。
   - **安装**:
     ```bash
     go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
     ```

8. **gorename**
   - **地址**: `golang.org/x/tools/cmd/gorename`
   - **简介**: `gorename` 用于全局重命名 Go 代码中的符号。
   - **安装**:
     ```bash
     go install -v golang.org/x/tools/cmd/gorename@latest
     ```

9. **goreturns**
   - **地址**: `github.com/sqs/goreturns`
   - **简介**: `goreturns` 格式化工具，自动填充函数的返回语句。
   - **安装**:
     ```bash
     go install -v github.com/sqs/goreturns@latest
     ```

10. **golangci-lint**
    - **地址**: `github.com/golangci/golangci-lint/cmd/golangci-lint`
    - **简介**: `golangci-lint` 是一个 Go 语言静态代码分析工具集合。
    - **安装**:
      ```bash
      go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
      ```

11. **guru**
    - **地址**: `golang.org/x/tools/cmd/guru`
    - **简介**: `guru` 提供 Go 代码分析和导航功能。
    - **安装**:
      ```bash
      go install -v golang.org/x/tools/cmd/guru@latest
      ```

12. **gomodifytags**
    - **地址**: `github.com/fatih/gomodifytags`
    - **简介**: `gomodifytags` 用于管理 Go 结构体的标签。
    - **安装**:
      ```bash
      go install -v github.com/fatih/gomodifytags@latest
      ```

13. **gotests**
    - **地址**: `github.com/cweill/gotests`
    - **简介**: `gotests` 根据现有 Go 代码自动生成测试代码。
    - **安装**:
      ```bash
      go install -v github.com/cweill/gotests/gotests@latest
      ```

14. **go-outline**
    - **地址**: `github.com/ramya-rao-a/go-outline`
    - **简介**: `go-outline` 生成 Go 代码文件的大纲。
    - **安装**:
      ```bash
      go install -v github.com/ramya-rao-a/go-outline@latest
      ```

15. **staticcheck**
    - **地址**: `honnef.co/go/tools/cmd/staticcheck`
    - **简介**: `staticcheck` 高级静态分析工具，用于检测 Go 代码中的问题。
    - **安装**:
      ```bash
      go install -v honnef.co/go/tools/cmd/staticcheck@latest
      ```

16. **impl**
    - **地址**: `github.com/josharian/impl`
    - **简介**: `impl` 生成接口的实现代码。
    - **安装**:
      ```bash
      go install -v github.com/josharian/impl@latest
      ```

17. **gopls (Go Language Server)**
    - **地址**: `golang.org/x/tools/gopls`
    - **简介**: `gopls` 实现了语言服务器协议 (LSP)，提供代码补全等功能。
    - **安装**:
      ```bash
      go install -v golang.org/x/tools/gopls@latest
      ```

## visualstudio 编辑工具使用
### 参考: https://learn.microsoft.com/zh-cn/azure/developer/go/configure-visual-studio-code
### 配置 CRLF 为 LF
    File -> Preference -> Settings
    搜索 files.eol
    修改 'Auto' 为 '\n'
    关闭窗口即可
  
### 打开 src 代码存放目录
    File -> Open Folder 
    选择目录 D:\GolandProjects\src\
    点击 选择文件夹 即可打开

### 安装 GO 环境插件
    左侧栏 -> Extensions(Ctrl+Shift+X)
    搜索 golang.go 点击 Install

### 安装运行插件
    左侧栏 -> Extensions(Ctrl+Shift+X)
    搜索 formulahendry.code-runner 点击 Install
    安装后直接打开 main.go 源码，光标右键点击 Run Code 即可运行代码

### 安装中文插件(可选)
    左侧栏 -> Extensions(Ctrl+Shift+X)
    搜索 vscode-language-pack-zh-hans 点击 Install
    安装后重启 visualstudio 可以看到效果

## goland idea 编辑工具配置
### 配置 Ctrl + 鼠标 缩放字体
    File -> Settings -> Editor -> General -> Mouse Control    
    勾选 Change font size with Ctrl+Mouse Wheel in:  

### 配置字体大小
    File -> Settings -> Editor -> Font  
    调整 Size

### 配置主题
    File -> Settings -> Editor -> Color Scheme  
    选择自定义主题 Scheme

### 添加自动化工具
    File -> Settings -> Tools -> File Watchers  
    添加自动格式化工具 go fmt  
    添加自动导入工具 goimports  

### 点击确定即可

