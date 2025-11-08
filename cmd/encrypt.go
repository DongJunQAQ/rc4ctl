package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"rc4ctl/internal"
)

func init() { // 将加密子命令添加到根命令
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{ //定义加密子命令
	Use:   "encrypt [input-file-path]",
	Short: "使用RC4算法加密文件",
	Long:  `将指定图片文件通过RC4算法加密，输出为不可直接查看的文件（需用decrypt子命令解密）`,
	Args:  cobra.ExactArgs(1), // 此子命令后面必须传入1个位置参数：输入要加密的文件路径
	PreRun: func(cmd *cobra.Command, args []string) { // Cobra提供的钩子函数，会在命令的Run函数执行前自动调用，即在执行前检查必需参数
		//cmd: 当前正在执行的命令对象（如encrypt或decrypt子命令）
		//args: 命令接收的位置参数（如输入文件路径）
		if keyFlag == "" || outputFlag == "" {
			cmd.Printf("错误：必须提供密钥和输出文件路径\n") //若某一参数为空，则打印错误提示
			err := cmd.Usage()               //打印当前命令的使用帮助
			if err != nil {
				return
			}
			os.Exit(1) //若校验参数失败则以状态码1退出程序
		}
	},
	Run: func(_ *cobra.Command, args []string) { // 加密逻辑执行函数，args的值为位置参数即需要加密的文件路径
		internal.ProcessCrypt(args, "加密", outputFlag, keyFlag)
	},
}
