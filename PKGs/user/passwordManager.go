package user

import (
	"io/ioutil"
	"liteDB/PKGs/variables"
	"fmt"
)


func ReadAllUsers() {
	content, err := ioutil.ReadFile(variables.PasswordManagerPath)
	HandleError(err)
	fmt.Println(string(content))
}