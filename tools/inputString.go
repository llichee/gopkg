package tools

import (
	"fmt"
	"strings"
)

func InputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)  // TrimSpace去掉 input字串 前后空格
}
