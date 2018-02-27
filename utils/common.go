package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

//md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	md5str := hex.EncodeToString(h.Sum(nil))
	return md5str
}

func NowInS() int64 {
	return time.Now().Unix()
}
