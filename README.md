# 创建项目

- go mod init <module_name>

# 相关命令

## 下载相关

### go get <module_name>

（1）依赖的下载位置：

- 默认行为：

  - go get 会将依赖项下载到模块缓存中，默认路径是 $GOPATH/pkg/mod（如果没有设置 GOMODCACHE，则默认为 ~/go/pkg/mod）。

  - 模块缓存是一个全局共享的存储区域，所有项目都可以从中复用已下载的模块，避免重复下载。

（2）是否直接放到 $GOPATH/pkg/mod？

- 是的，go get 默认会将依赖项放入模块缓存（即 $GOPATH/pkg/mod 或 GOMODCACHE 指定的目录）。

- 同时，go get 还会更新当前项目的 go.mod 和 go.sum 文件，以记录新添加或更新的依赖项及其版本。

（3）编译行为

- 在 Go 1.16 及更高版本 中，go get 不再编译代码或生成可执行文件。它的主要作用是管理依赖项 （添加、更新或移除模块）。

- 如果你需要编译并生成可执行文件，请使用 go install 或 go build。默认会将依赖项下载到模块缓存中，而不是直接放到 $GOPATH/pkg/mod 目录中

### go install <module_name>

（1）拉取项目

- go install 会从远程仓库拉取指定的模块，并将其编译成可执行文件。

- 如果模块包含多个包，go install 会尝试编译模块的主包（即包含 main 函数的包）。

（2）编译后的可执行文件位置

- 编译后的可执行文件会被放入 $GOPATH/bin 或 $GOBIN 指定的目录中：

  - 如果未设置 GOBIN，默认路径是 $GOPATH/bin。

  - 如果未设置 GOPATH，默认路径是 ~/go/bin。

- 示例：

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

  这会在 $GOPATH/bin 或 $GOBIN 中生成一个名为 golangci-lint 的可执行文件。

（3）与 go get 的区别

在 Go 1.16 及更高版本 中，go install 和 go get 的职责被明确分开：

- go get：用于管理依赖项（更新 go.mod 和 go.sum），不再生成可执行文件。

- go install：用于安装工具或应用程序，生成可执行文件。

3. 总结

- go get <module_name> ：

  - 将依赖项下载到模块缓存（默认 $GOPATH/pkg/mod 或 GOMODCACHE）。

  - 更新 go.mod 和 go.sum 文件。

  - 不生成可执行文件 （Go 1.16+）。

- go install <module_name> ：

  - 拉取项目并编译生成可执行文件。

  - 将编译后的可执行文件放入 $GOPATH/bin 或 $GOBIN。

## 查看依赖

- go list -m all

  这将列出所有模块及其版本。如果你想确保依赖项被下载到本地，可以使用 go mod download 命令

- go mod download

  下载并将所有依赖项拷贝到本地的 `$GOPATH/pkg/mod` 目录

- ls $(go env GOPATH)/pkg/mod

  查看已经下载的三方项目

# 测试

[testify](https://github.com/stretchr/testify) 具有常见断言和模拟的工具包，可以与标准库很好地配合

# 命令行

[cobra](https://github.com/spf13/cobra)

[color](https://github.com/fatih/color)
