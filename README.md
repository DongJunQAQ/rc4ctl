# rc4ctl
一个基于 Go 语言和 RC4 对称加密算法开发的命令行工具（rc4ctl），支持通过 `encrypt` 和 `decrypt` 子命令对文件进行加解密操作，需指定输入文件路径、输出文件路径（-o/--output）和密钥（-k/--key），适用于对图片等文件进行简单的加解密操作。

## 构建：

编译项目生成二进制文件：

```shell
make
```

清除编译产物：

```shell
make clean
```



## 快速开始：

- 加密图片：

```
.\rc4img.exe encrypt .\image\demo.png -o one.bin -k 123
```

- 解密图片：

```
.\rc4img.exe decrypt .\one.bin -o two.png -k 123
```



## 使用指南：

```shell
PS E:\rc4img> .\rc4img.exe -h                                 
基于RC4对称加密算法，对JPG/PNG/GIF等图片文件进行加密和解密，加密解密使用相同密钥

Usage:
  rc4img [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     解密RC4加密的文件
  encrypt     加密图片文件
  help        Help about any command

Flags:
  -h, --help            help for rc4img
  -k, --key string      加解密密钥（必填）
  -o, --output string   输出文件路径（必填）

Use "rc4img [command] --help" for more information about a command.
```

子命令用法：

```shell
PS E:\rc4img> .\rc4img.exe encrypt -h
将指定图片文件通过RC4算法加密，输出为不可直接查看的文件（需用decrypt子命令解密）

Usage:
  rc4img encrypt [input-file] [flags]

Flags:
  -h, --help   help for encrypt

Global Flags:
  -k, --key string      加解密密钥（必填）
  -o, --output string   输出文件路径（必填）
PS E:\rc4img>
```

## 命令补全：

```shell

```

