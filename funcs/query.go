package funcs

import (
	_ "fmt"
	"strings"
	"gocode/tools"
)

func Query(users map[int]map[string]string) {
	q := tools.InputString("please input query connect:")
	for k, v := range users {
		// id, name, birththday, tel, addr, desc
		if strings.Contains(v["name"], q) || strings.Contains(v["birthday"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["addr"], q) {
			tools.PrintUser(k, v)
		}
	}
}
