package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	passwd := "af33208d995a60c8c0c05dc6caca593d"
	fmt.Println("please input password:")
	var password string
	fmt.Scan(&password)
	var p = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if passwd == p {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
