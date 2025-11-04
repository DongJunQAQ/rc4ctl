package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"rc4img/cmd"
)

// encryptCmd 定义加密子命令
var encryptCmd = &cobra.Command{
	Use:   "encrypt [input-file]",
	Short: "加密图片文件",
	Long:  `将指定图片文件通过RC4算法加密，输出为不可直接查看的文件（需用decrypt子命令解密）`,
	Args:  cobra.ExactArgs(1), // 必须传入1个位置参数：输入文件路径
	Run:   runEncrypt,         // 加密逻辑执行函数
}

func init() {
	cmd.RootCmd.AddCommand(encryptCmd) // 将加密子命令添加到根命令
}

// runEncrypt 执行加密逻辑
func runEncrypt(_ *cobra.Command, args []string) {
	inputPath := args[0] // 位置参数：输入图片路径
	outputPath := cmd.OutputFlag
	key := cmd.KeyFlag
	// 读取输入文件
	data, err := readFile(inputPath)
	if err != nil {
		fmt.Printf("加密失败：%v\n", err)
		return
	}
	// RC4加密
	encryptedData := rc4Crypt(data, []byte(key))
	// 写入输出文件
	if err := writeFile(outputPath, encryptedData); err != nil {
		fmt.Printf("加密失败：%v\n", err)
		return
	}
	fmt.Printf("加密成功！\n输入：%s\n输出：%s\n", inputPath, outputPath)
}
