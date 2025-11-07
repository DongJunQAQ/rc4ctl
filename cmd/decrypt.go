package cmd

import (
	"github.com/spf13/cobra"
	"rc4ctl/internal"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
	// 标记此子命令的flag参数为必填，未传则报错，根命令不受影响
	_ = rootCmd.MarkPersistentFlagRequired("key")
	_ = rootCmd.MarkPersistentFlagRequired("output")
}

var decryptCmd = &cobra.Command{ //定义解密子命令
	Use:   "decrypt [input-file-path]",
	Short: "使用RC4算法解密文件",
	Long:  `将通过encrypt子命令加密的文件解密，恢复为原始可查看的图片文件（需使用加密时的相同密钥）`,
	Args:  cobra.ExactArgs(1), // 此子命令后面必须传入1个位置参数：输入要解密的文件路径
	Run: func(_ *cobra.Command, args []string) { // 解密逻辑执行函数，args的值为位置参数即需要解密的文件路径
		internal.ProcessCrypt(args, "解密", outputFlag, keyFlag)
	},
}
