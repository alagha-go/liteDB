package user

import (
	"io/ioutil"
	"liteDB/PKGs/variables"
)


func ReadAllUsers() string {
	content, err := ioutil.ReadFile(variables.PasswordManagerPath)
	HandleError(err)

	return string(content)
}