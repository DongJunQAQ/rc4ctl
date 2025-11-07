# rc4ctl
一个基于 Go 语言和 RC4 对称加密算法开发的命令行工具（rc4ctl），支持通过 `encrypt` 和 `decrypt` 子命令对文件进行加解密操作，需指定输入文件路径、输出文件路径（-o/--output）和密钥（-k/--key），适用于对图片等文件进行简单的加解密操作。

## 构建：

- 编译项目生成二进制文件：

```shell
make
```

执行完此命令后会在./bin目录下生成名为rc4ctl的可执行文件；

- 验证是否安装成功：

```
[root@ecs-1622 rc4ctl-master]# cd ./bin/ && ./rc4ctr -v
rc4ctl version 0.1.0
```

- 清除编译产物：

```shell
make clean
```



## 快速开始：

- 加密图片：

```
[root@ecs-1622 bin]# ./rc4ctr encrypt ../image/demo.jpg -o cipher.bin -k 1234
加密成功！
输入：../image/demo.jpg
输出：cipher.bin
```

- 解密图片：

```
[root@ecs-1622 bin]# ./rc4ctr decrypt cipher.bin -o plaintext.jpg -k 1234
解密成功！
输入：cipher.bin
输出：plaintext.jpg
```



## 使用指南：

```shell
[root@ecs-1622 bin]# ./rc4ctr -h
基于RC4对称加密算法，对JPG/PNG/GIF等图片文件进行加密和解密，加密解密使用相同密钥

Usage:
  rc4ctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     使用RC4算法解密文件
  encrypt     使用RC4算法加密文件
  help        Help about any command

Flags:
  -h, --help            help for rc4ctl
  -k, --key string      加解密密钥（必填）
  -o, --output string   输出文件路径（必填）
  -v, --version         version for rc4ctl

Use "rc4ctl [command] --help" for more information about a command.
```

子命令用法：

```shell
[root@ecs-1622 bin]# ./rc4ctr encrypt -h
将指定图片文件通过RC4算法加密，输出为不可直接查看的文件（需用decrypt子命令解密）

Usage:
  rc4ctl encrypt [input-file-path] [flags]

Flags:
  -h, --help   help for encrypt

Global Flags:
  -k, --key string      加解密密钥（必填）
  -o, --output string   输出文件路径（必填）
```

## 命令补全：

```shell

```

