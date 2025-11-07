package internal

import (
	"fmt"
	"io"
	"os"
)

func readFile(path string) ([]byte, error) { //读取文件内容
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件：%s", path)
	}
	defer file.Close()
	data, err := io.ReadAll(file) //从打开的文件中读取所有数据，直到文件结束
	if err != nil {
		return nil, fmt.Errorf("读取文件失败：%s", path)
	}
	return data, nil
}

func writeFile(path string, data []byte) error { //写入文件内容
	file, err := os.Create(path) //创建指定路径的文件，若文件已存在则截断（清空内容）
	if err != nil {
		return fmt.Errorf("无法创建文件：%s", path)
	}
	defer file.Close()
	_, err = file.Write(data) //将字节数据写入文件
	if err != nil {
		return fmt.Errorf("写入文件失败：%s", path)
	}
	return nil
}
