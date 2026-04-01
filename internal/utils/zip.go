package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

// 将字符串进行 gzip 压缩，并返回 base64 编码的结果
// 参数 data 是要压缩的字符串
// 返回值是 base64 编码的压缩字符串
func Gzip(data string) (string, error) {
	var buf bytes.Buffer
	// 使用最佳压缩级别，与 Python 的 zlib.Z_BEST_COMPRESSION 对应
	writer, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return "", fmt.Errorf("创建 gzip 写入器失败：%w", err)
	}
	if _, err := writer.Write([]byte(data)); err != nil {
		return "", fmt.Errorf("压缩过程中出错：%w", err)
	}
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("关闭压缩写入器时出错：%w", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// 解码 base64 字符串，进行 gzip 解压缩，并返回 UTF-8 字符串
// 参数 data 是要解压缩的 base64 编码字符串
// 返回值是解压缩后的原始字符串
func UnGzip(data string) (string, error) {
	// 解码 base64 字符串
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("base64 解码失败：%w", err)
	}
	// 使用 gzip.NewReader 进行解压缩，自动检测 gzip 头部
	reader, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return "", fmt.Errorf("创建 gzip 读取器失败：%w", err)
	}
	defer reader.Close()
	result, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("读取压缩数据失败：%w", err)
	}
	return string(result), nil
}
