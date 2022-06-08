package zincapi

import "strings"

// addressShaping 连接地址整形
func addressShaping(addr string) string {
	if strings.HasSuffix(addr, "/") {
		return addr[:len(addr)-1]
	}
	return addr
}
