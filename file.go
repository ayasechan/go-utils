package goutils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var CompliedUnsafeChar = regexp.MustCompile(`[<>:"\\/\|\?\*]`)

// 移除win下不可用于文件名的字符
func AsSafeFileName(s string, to string) string {
	return strings.TrimSpace(CompliedUnsafeChar.ReplaceAllString(s, to))
}

// The final path component, without its suffix
func GetFileStem(src string) string {
	return strings.TrimSuffix(filepath.Base((src)), filepath.Ext(src))
}

// Return a new path with the stem changed.
func SetFileStem(src, stem string) string {
	ext := filepath.Ext(src)
	return SetFileName(src, fmt.Sprintf("%s%s", stem, ext))
}

// Return a new path with the suffix changed.
// If the original path doesn’t have a suffix, the new suffix is appended instead.
// If the suffix is an empty string, the original suffix is removed.
// suffix should start with dot!
func SetFileSuffix(src, suffix string) string {
	ext := filepath.Ext(src)
	if ext == suffix {
		return filepath.Clean(src)
	}
	if ext == "" {
		return filepath.Clean(fmt.Sprintf("%s%s", src, suffix))
	}
	return SetFileName(src, fmt.Sprintf("%s%s", GetFileStem(src), suffix))
}

// Return a new path with the name changed.
func SetFileName(src, name string) string {
	return filepath.Join(filepath.Dir(src), name)
}
