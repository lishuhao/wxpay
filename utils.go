package wxpay

import (
	"strconv"
	"time"
)

// 生成随机字符串
func nonceStr() string {
	return strconv.FormatInt(time.Now().Unix(),10)
}
