package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Encryption(str string) string {
	w := md5.New()
	io.WriteString(w, str) //将str写入到w中
	bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式

	// md5str2 := fmt.Sprintf("%x", bw)    //将 bw 转成字符串
	md5str := hex.EncodeToString(bw) //将 bw 转成字符串
	return md5str
}
