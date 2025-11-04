package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"rc4img/cmd"
)

// decryptCmd 定义解密子命令
var decryptCmd = &cobra.Command{
	Use:   "decrypt [input-file]",
	Short: "解密RC4加密的文件",
	Long:  `将通过encrypt子命令加密的文件解密，恢复为原始可查看的图片文件（需使用加密时的相同密钥）`,
	Args:  cobra.ExactArgs(1), // 必须传入1个位置参数：加密文件路径
	Run:   runDecrypt,         // 解密逻辑执行函数
}

func init() {
	cmd.RootCmd.AddCommand(decryptCmd) // 将解密子命令添加到根命令
}

// runDecrypt 执行解密逻辑（RC4加密和解密逻辑完全一致）
func runDecrypt(_ *cobra.Command, args []string) {
	inputPath := args[0] // 位置参数：加密文件路径
	outputPath := cmd.OutputFlag
	key := cmd.KeyFlag
	// 读取加密文件
	data, err := readFile(inputPath)
	if err != nil {
		fmt.Printf("解密失败：%v\n", err)
		return
	}
	// RC4解密（与加密共用函数）
	decryptedData := rc4Crypt(data, []byte(key))
	// 写入输出文件（恢复为图片）
	if err := writeFile(outputPath, decryptedData); err != nil {
		fmt.Printf("解密失败：%v\n", err)
		return
	}
	fmt.Printf("解密成功！\n输入：%s\n输出：%s\n", inputPath, outputPath)
}
