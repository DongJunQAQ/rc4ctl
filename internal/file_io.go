package internal

import (
	"fmt"
	"io"
	"os"
)

// readFile 读取文件内容
func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件：%s", path)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败：%s", path)
	}
	return data, nil
}

// writeFile 写入文件内容
func writeFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("无法创建文件：%s", path)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("写入文件失败：%s", path)
	}
	return nil
}
