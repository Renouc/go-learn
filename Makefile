# TODO 将启动web服务的命令写在这里

# 定义变量
APP_NAME=gold

# 默认目标
all: help

# 启动 Web 服务
web: 
	go run main.go web

# 📦 打包
build:
	go build -o $(APP_NAME) main.go

# 安装可执行文件

# 将当前项目编译为可执行文件，放入 go/bin 路径下
# 默认会将 go.mod 中的 module 作为可执行文件名
install:
	go install .

# 卸载可执行文件
uninstall:
	rm `go env GOPATH`/bin/${APP_NAME}

# 查看bin
ls:
	ls `go env GOPATH`/bin

# 显示帮助信息
help:
	@echo "Available commands:"
	@echo "  web        - Start the web server"
	@echo "  build      - Build the application"
	@echo "  install    - Install the application"
	@echo "  uninstall  - Uninstall tha application"
	@echo "  ls         - List executables in the GOPATH/bin directory"