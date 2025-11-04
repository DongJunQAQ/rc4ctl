package main

import (
	"rc4img/cmd"
	_ "rc4img/internal" // 只有导入internal包后才能触发init函数
)

func main() {
	cmd.Execute()
}
