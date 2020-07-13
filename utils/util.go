package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// GetVideoName 生成视频文件名
func GetVideoName(uid string) string {
	h := md5.New()
	h.Write([]byte(uid + strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hex.EncodeToString(h.Sum(nil))
}
