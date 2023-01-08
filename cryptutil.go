package scalpel

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func HashByMd5(s string, toUpper bool) string {
	has := md5.Sum([]byte(s))
	if toUpper {
		return strings.ToUpper(fmt.Sprintf("%x", has))
	}
	return fmt.Sprintf("%x", has) // 将[]byte转成16进制

}
