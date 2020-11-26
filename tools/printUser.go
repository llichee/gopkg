package tools

import "fmt"

func PrintUser(pk int, user map[string]string) {
	fmt.Println("id:", pk)
	fmt.Println("name:", user["name"])
	fmt.Println("birthday:", user["birthday"])
	fmt.Println("tel:", user["tel"])
	fmt.Println("addr:", user["addr"])
	fmt.Println("desc:", user["desc"])
}
