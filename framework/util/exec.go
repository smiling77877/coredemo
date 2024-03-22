package util

import "os"

// GetExecDirectory获取当前执行程序目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return file + "/"
	}
	return ""
}
