package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// 自定义错误
var (
	ErrInvalidPKCS7Padding = errors.New("invalid PKCS7 padding")
)

// pkcs7Pad 对数据进行 PKCS7 填充
// blockSize 通常为 16 (AES)
func PKCS7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad 移除 PKCS7 填充
// 返回去填充后的数据和错误
func PKCS7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	padding := int(data[length-1])
	if padding < 1 || padding > 32 {
		return nil, ErrInvalidPKCS7Padding
	}

	// 验证所有填充字节是否一致
	for i := length - padding; i < length; i++ {
		if data[i] != byte(padding) {
			return nil, ErrInvalidPKCS7Padding
		}
	}

	return data[:length-padding], nil
}

// AESEncrypt 使用AES-CBC模式对明文进行加密并返回Base64编码的密文
// 此函数实现了AES加密功能，使用PKCS7填充和CBC加密模式，最后将密文进行Base64编码
//
// 参数:
//
//	plainText: 要加密的明文字符串
//	keyStr: 加密使用的密钥字符串
//
// 返回值:
//
//	string: 加密后的Base64编码字符串
//	error: 加密过程中遇到的错误，如密钥无效等
func AESEncrypt(plainText, keyStr string) (string, error) {
	// 将密钥字符串转换为字节数组
	key := []byte(keyStr)

	// 创建AES加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 1. 对明文进行PKCS7填充，确保数据长度为AES块大小的整数倍
	padded := PKCS7Pad([]byte(plainText), aes.BlockSize)

	// 2. 分配密文存储空间，长度与填充后的明文相同
	encrypted := make([]byte, len(padded))

	// 3. 使用密钥的前16字节作为初始化向量(IV)
	// 注意：这种做法与Python版本保持一致，但在实际应用中，为了更高的安全性，
	// 通常应使用随机生成的IV并随密文一起传输
	iv := key[:aes.BlockSize]

	// 4. 创建CBC模式加密器并执行加密操作
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, padded)

	// 5. 对加密后的字节数组进行Base64编码，转换为可传输的字符串格式
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// AESDecrypt 使用AES-CBC模式解密Base64编码的加密字符串
// 此函数实现了与AESEncrypt相对应的解密功能，使用相同的密钥和初始化向量
//
// 参数:
//
//	encryptedBase64: 经过Base64编码的加密字符串
//	keyStr: 解密使用的密钥字符串
//
// 返回值:
//
//	string: 解密后的原始明文
//	error: 解密过程中遇到的错误，包括Base64解码错误、密钥无效、密文格式错误等
func AESDecrypt(encryptedBase64, keyStr string) (string, error) {
	// 将密钥字符串转换为字节数组
	key := []byte(keyStr)

	// 1. 对输入的Base64编码字符串进行解码
	data, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	// 2. 创建AES加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 3. 验证密文长度是否有效（必须为AES块大小的整数倍）
	if len(data) == 0 || len(data)%aes.BlockSize != 0 {
		return "", errors.New("invalid ciphertext block size")
	}

	// 4. 分配解密数据的缓冲区
	decrypted := make([]byte, len(data))

	// 5. 使用密钥的前16字节作为初始化向量（与Python版本保持一致）
	iv := key[:aes.BlockSize]

	// 6. 创建CBC模式解密器并执行解密操作
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, data)

	// 7. 去除PKCS7填充，获取原始明文
	plainText, err := PKCS7Unpad(decrypted)
	if err != nil {
		return "", err
	}

	// 8. 将解密后的字节数组转换为字符串并返回
	return string(plainText), nil
}

// Base64Encode 将 []byte 编码为 Base64 字符串
// 输入：原始字节
// 输出：Base64 编码后的字符串
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode 将 Base64 字符串解码为原始字符串
// 输入：Base64 字符串
// 输出：解码后的字符串 + 错误
func Base64Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("Base64 解码失败: %w", err)
	}
	return string(decoded), nil
}

// MD5Hash 计算字符串的 MD5 哈希值（小写，32位）
// str: 要计算哈希的字符串
// 返回 MD5 哈希值（32位小写十六进制字符串）和错误
func MD5Hash(str string) (string, error) {
	// 计算 MD5 哈希
	hash := md5.Sum([]byte(str))

	// 转换为十六进制字符串
	hashStr := hex.EncodeToString(hash[:])

	// 确保是 32 位（MD5 应该是 32 位，但为了安全起见）
	if len(hashStr) < 32 {
		// 如果不足 32 位，前面补 0
		hashStr = fmt.Sprintf("%032s", hashStr)
	}

	return hashStr, nil
}
