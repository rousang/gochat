package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写的md5加密
func Md5Encode(data string) string {
	e := md5.New()
	e.Write([]byte(data))
	return hex.EncodeToString(e.Sum(nil))
}

// 大写的md5加密
func Md5EncodeUpper(data string) string {
	return strings.ToUpper(Md5Encode(data))
}
