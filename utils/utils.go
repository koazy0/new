package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// MD5 返回十六进制字符串
func MD5(bytedata []byte) string {
	hash := md5.New()
	hash.Write(bytedata)                      //写入数据到哈希对象
	md5Bytes := hash.Sum(nil)                 //计算哈希值
	md5String := hex.EncodeToString(md5Bytes) //将字节切片转换为十六进制字符串
	return md5String
}

// GetExtend 获取文件小写扩展名
func GetExtend(str string) string {

	strs := strings.Split(str, ".")
	if len(strs) <= 1 {
		return ""
	}
	return strings.ToLower(strs[len(strs)-1])
}

// CheckWhiteImageList 对图片白名单进行校验
// 如果在白名单中（允许通过）则返回true
func CheckWhiteImageList(key string, list []string) bool {
	return InList(GetExtend(key), list)
}
