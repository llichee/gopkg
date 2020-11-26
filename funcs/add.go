package funcs

import (
	"gocode/tools"
)

func Add(users map[int]map[string]string) {
	id := tools.GetId(users)
	users[id] = tools.InputUser()
}
