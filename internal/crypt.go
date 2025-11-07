package internal

import "fmt"

func ProcessCrypt(args []string, operation string, outputPath string, key string) { // 加/解密共用的处理函数
	inputPath := args[0] //要加/解密的文件路径
	// 读取文件
	data, err := readFile(inputPath)
	if err != nil {
		fmt.Printf("%s失败：%v\n", operation, err)
		return
	}
	// 执行RC4加解密（本身算法对称）
	processedData := rc4Crypt(data, []byte(key)) //将读取到的文件内容和加/解密所使用的key传递给RC4算法，让其加/解密并获取其处理后的内容
	// 写入输出文件
	if err := writeFile(outputPath, processedData); err != nil { //将-o选项后接的文件路径与加/解密后的内容输出为一个新文件
		fmt.Printf("%s失败：%v\n", operation, err)
		return
	}
	fmt.Printf("%s成功！\n输入：%s\n输出：%s\n", operation, inputPath, outputPath)
}
