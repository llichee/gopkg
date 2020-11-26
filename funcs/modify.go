package funcs

import (
	"fmt"
	"strconv"
	"gocode/tools"
)

func Modify(users map[int]map[string]string) {
	idString := tools.InputString("please input modify user id:")
	if id, err := strconv.Atoi(idString); err == nil {
		if _, ok := users[id]; ok {
			fmt.Println("将修改的用户信息：")
			fmt.Println(users)
			input := tools.InputString("确定修改(Y/N)?")
			if input == "y" || input == "Y" {
				user := tools.InputUser()
				users[id] = user
				fmt.Println("修改成功")
			}
		}else {
			fmt.Println("ID is noexit")
		}
	}else {
		fmt.Println("ID is error!")
	}
}
