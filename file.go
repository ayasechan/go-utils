package goutils

import (
	"regexp"
	"strings"
)

var CompliedUnsafeChar = regexp.MustCompile(`[<>:"\\/\|\?\*]`)

// 移除win下不可用于文件名的字符
func AsSafeFileName(s string, to string) string {
	return strings.TrimSpace(CompliedUnsafeChar.ReplaceAllString(s, to))
}
