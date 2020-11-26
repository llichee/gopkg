package tools

func InputUser() map[string]string {
	user := map[string]string{}
	user["name"] = InputString("please input name:")
	user["birthday"] = InputString("please input birthday(2000-01-01):")
	user["tel"] = InputString("please input tel:")
	user["addr"] = InputString("please input addr:")
	user["desc"] = InputString("please input desc:")
	return user
}
