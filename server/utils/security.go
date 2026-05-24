package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func AesGCMEncrypt(plainText []byte, key []byte) (string, error) {
	// 1. 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 2. 创建 GCM 模式实例
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 3. 生成随机 IV (GCM 标准推荐 12 字节)
	iv := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	// 4. 加密：cipherText = IV + 密文(含认证标签)
	cipherText := gcm.Seal(iv, iv, plainText, nil)

	// 5. 返回 Base64 编码字符串（方便传输/存储）
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AesGCMDecrypt(cipherTextBase64 string, key []byte) ([]byte, error) {
	// 1. 解码 Base64
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return nil, err
	}

	// 2. 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 3. 创建 GCM 模式实例
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 4. 拆分 IV 和 密文
	ivSize := gcm.NonceSize()
	iv := cipherText[:ivSize]
	realCipherText := cipherText[ivSize:]

	// 5. 解密（自动校验数据完整性，被篡改会直接报错）
	plainText, err := gcm.Open(nil, iv, realCipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
