package internal

func rc4Crypt(data []byte, K []byte) []byte { //RC4加解密核心函数（加密和解密逻辑一致）
	S := make([]int, 256)
	T := make([]int, 256)
	keyLen := len(K)
	// 初始化S盒和T盒
	for i := 0; i <= 255; i++ {
		S[i] = i
		T[i] = int(K[i%keyLen])
	}
	// 置换S盒
	j := 0
	for i := 0; i <= 255; i++ {
		j = (j + S[i] + T[i]) % 256
		S[i], S[j] = S[j], S[i] //交换二者的值
	}
	// 生成密钥流并异或
	result := make([]byte, len(data))
	i, j := 0, 0
	for k := 0; k < len(data); k++ {
		i = (i + 1) % 256
		j = (j + S[i]) % 256
		S[i], S[j] = S[j], S[i]
		t := (S[i] + S[j]) % 256
		result[k] = data[k] ^ byte(S[t]) //^表示按位异或运算
	}
	return result
}
