package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var ( // 定义全局参数即所有子命令共用的「密钥」和「输出文件路径」
	keyFlag    string // -k/--key 密钥
	outputFlag string // -o/--output 输出文件路径
)

func init() {
	// 为根命令绑定全局flag参数（所有子命令自动继承），给rootCmd添加两个持久化字符串参数
	rootCmd.PersistentFlags().StringVarP(&keyFlag, "key", "k", "", "加解密密钥（必填）") //StringVarP定义的是flag后面接的值的类型
	//如果是BoolVarP()类型则用户使用时无需传值，只需指定参数本身即可表示true，如--debug等价于debugFlag=true；若命令中不指定此flag则为默认值（通常为false）
	rootCmd.PersistentFlags().StringVarP(&outputFlag, "output", "o", "", "输出文件路径（必填）")
}

var rootCmd = &cobra.Command{ // RootCmd 是所有子命令的基础根命令
	Use:     "rc4ctl",                                          // 命令名称，在终端中使用时的指令（如./rc4img）
	Short:   "RC4算法图片加解密工具",                                    // 短描述
	Long:    `基于RC4对称加密算法，对JPG/PNG/GIF等图片文件进行加密和解密，加密解密使用相同密钥`, //长描述
	Version: "0.1.0",                                           //添加此工具的版本信息
}

func Execute() { //初始化并执行根命令
	err := rootCmd.Execute() // 执行根命令，会解析命令行参数并触发对应逻辑
	if err != nil {          // 如果执行过程中发生错误
		os.Exit(1) // 退出程序，返回状态码1（表示执行失败）
	}
}

/*
该项目是如何获取到flag后面的值并应用到整条命令中去的？
1.用户输入命令（如rc4ctl encrypt input.png -o output.bin -k 123）。
2.cobra解析命令行参数，将-k 123赋值给变量keyFlag，-o output.bin赋值给变量outputFlag。
3.调用子命令的Run函数，传入解析后的flag值到业务逻辑函数（internal.ProcessCrypt）。
4.最终在ProcessCrypt中使用密钥和输出路径完成加解密操作。
*/
