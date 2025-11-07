# 定义项目名称
BINARY_NAME := rc4ctr
# 定义编译输出目录
BUILD_DIR := ./bin
# 定义源代码入口
MAIN_PACKAGE := ./main.go
# 默认行为，执行make时等同于执行make build
.DEFAULT_GOAL := build

# 根据不同的平台设置不同的变量
ifeq ($(OS),Windows_NT)  # 如果当前系统为Windows_NT的话执行一下操作
    BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME).exe
    RMDIR := powershell -Command "Remove-Item -Path $(BUILD_DIR) -Recurse -Force"
else
    BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)
    RMDIR := rm -rf $(BUILD_DIR)
endif

## build: 编译项目
build:
	@echo "开始在$(OS)平台上编译此项目..."
	go build -o $(BINARY_PATH) $(MAIN_PACKAGE)
	@echo "编译完成：$(BINARY_PATH)"

## clean: 清理编译产物
clean:
	@echo "清理编译产物..."
	$(RMDIR)
	@echo "清理完成"

## help: 展示帮助信息
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"