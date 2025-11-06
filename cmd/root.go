package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// 全局参数：所有子命令共用的「密钥」和「输出文件路径」
var (
	keyFlag    string // -k/--key 密钥
	outputFlag string // -o/--output 输出文件路径
)

func init() {
	// 为根命令绑定全局参数（所有子命令自动继承）
	RootCmd.PersistentFlags().StringVarP(&keyFlag, "key", "k", "", "加解密密钥（必填）")
	RootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "", "输出文件路径（必填）")
	// 标记参数为必填，未传则报错
	_ = RootCmd.MarkPersistentFlagRequired("key")
	_ = RootCmd.MarkPersistentFlagRequired("output")
}

// RootCmd 是所有子命令的基础根命令
var RootCmd = &cobra.Command{
	Use:   "rc4img",
	Short: "RC4算法图片加解密工具",
	Long:  `基于RC4对称加密算法，对JPG/PNG/GIF等图片文件进行加密和解密，加密解密使用相同密钥`,
}

// Execute 初始化并执行根命令
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
