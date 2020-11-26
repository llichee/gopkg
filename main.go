package main

import (
	"fmt"
	"gocode/funcs"
	"gocode/tools"
)


//const (
//	MaxAuth int = 3
//	Password = "hengda"
//)

//func inputString(prompt string) string {
//	var input string
//	fmt.Print(prompt)
//	fmt.Scan(&input)
//	return strings.TrimSpace(input)  // TrimSpace去掉 input字串 前后空格
//}

//func auth() bool {
//	var input string
//	for i := 0; i <= maxAuth; i++ {
//		fmt.Println("please input password:")
//		fmt.Scan(&input)
//		if password == input {
//			return true
//		} else {
//			fmt.Println("password is false!!!")
//		}
//	}
//	return false
//}


//func query(users map[int]map[string]string) {
//	q := inputString("please input query connect:")
//	for k, v := range users {
//		// id, name, birththday, tel, addr, desc
//		if strings.Contains(v["name"], q) || strings.Contains(v["birthday"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["addr"], q) {
//			fmt.Println(k, v)
//		}
//	}
//}

//func getId(users map[int]map[string]string) int {
//	var id int
//	for k :=range users {
//		if id< k {
//			id = k
//		}
//	}
//	return id + 1
//}

//func add(users map[int]map[string]string) {
//	id := getId(users)
//	users[id] = inputUser()
//}

//func inputUser() map[string]string {
//	user := map[string]string{}
//	user["name"] = inputString("please input name:")
//	user["birthday"] = inputString("please input birthday(2000-01-01):")
//	user["tel"] = inputString("please input tel:")
//	user["addr"] = inputString("please input addr:")
//	user["desc"] = inputString("please input desc:")
//	return user
//}

//func modify(users map[int]map[string]string) {
//	idString := inputString("please input modify user id:")
//	if id, err := strconv.Atoi(idString); err == nil {
//		if _, ok := users[id]; ok {
//			fmt.Println("将修改的用户信息：")
//			fmt.Println(users)
//			input := inputString("确定修改(Y/N)?")
//			if input == "y" || input == "Y" {
//				user := inputUser()
//				users[id] = user
//				fmt.Println("修改成功")
//			}
//		}else {
//			fmt.Println("ID is noexit")
//		}
//	}else {
//		fmt.Println("ID is error!")
//	}
//}

//func del(users map[int]map[string]string) {
//	if id, err := strconv.Atoi(inputString("please input modify user id:")); err == nil {
//		if _, ok := users[id]; ok {
//			fmt.Println("将修改的用户信息：")
//			fmt.Println(users)
//			input := inputString("确定修改(Y/N)?")
//			if input == "y" || input == "Y" {
//				delete(users, id)
//				fmt.Println("[+]删除成功")
//			}
//		}else {
//			fmt.Println("ID is noexit")
//		}
//	}else {
//		fmt.Println("ID is error!")
//	}
//}

func main() {
	if !tools.Auth() {
		fmt.Printf("Input password %d false, exit\n", tools.MaxAuth)
		return
	}
	menu := `
++++++++++++++++
1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
++++++++++++++++
`
	// id, name, birththday, tel, addr, desc
	//users := map[int][5]string
	//users := map[int][]string
	//users := []map[string]string{}
	users := map[int]map[string]string{}
	callbacks := map[string]func(map[int]map[string]string) {
		"1": funcs.Query,
		"2": funcs.Add,
		"3": funcs.Modify,
		"4": funcs.Del,
	}

	fmt.Println("welcome my user manager os!")
//END:
	for {
		fmt.Println(menu)
		op := tools.InputString("please input your select:")
		callback, ok := callbacks[op]
		if ok {
			callback(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("指令错误")
		}
		//switch inputString("please input your select:") {
		//case "1":
		//	query(users)
		//case "2":
		//	add(users)
		//case "3":
		//	modify(users)
		//case "4":
		//	del(users)
		//case "5":
		//	break END
		//default:
		//	fmt.Println("指令错误！")
		//}
	}
}
