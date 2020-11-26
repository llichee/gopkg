package funcs

import (
	"fmt"
	"strconv"
	"gocode/tools"
)

func Del(users map[int]map[string]string) {
	if id, err := strconv.Atoi(tools.InputString("please input modify user id:")); err == nil {
		if _, ok := users[id]; ok {
			fmt.Println("将修改的用户信息：")
			fmt.Println(users)
			input := tools.InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("[+]删除成功")
			}
		}else {
			fmt.Println("ID is noexit")
		}
	}else {
		fmt.Println("ID is error!")
	}
}
