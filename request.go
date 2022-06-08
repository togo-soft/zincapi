package zincapi

import (
	"fmt"

	"github.com/guonaihong/gout"
)

func getHeader() gout.H {
	return gout.H{
		"Authorization": fmt.Sprintf("Basic %s", std.Token),
		"User-Agent":    "zincapi/0.1",
	}
}

func getURI(src string) string {
	return std.Address + src
}
