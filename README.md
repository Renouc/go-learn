# 创建项目

- go mod init <module_name>

# 相关命令

- go install <module_name>

  拉取项目，将编译后的可执行文件放入 go/bin 目录

- go get <module_name>

  默认会将依赖项下载到模块缓存中，而不是直接放到 $GOPATH/pkg/mod 目录中

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
