package tools

import (
	"crypto/md5"
	"fmt"
)

const (
	MaxAuth int = 3
	Password = "af33208d995a60c8c0c05dc6caca593d"
)

func Auth() bool {
	var input string
	for i := 0; i <= MaxAuth; i++ {
		fmt.Println("please input password:")
		fmt.Scan(&input)
		var p = fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if Password == p {
			return true
		} else {
			fmt.Println("password is false!!!")
		}
	}
	return false
}
