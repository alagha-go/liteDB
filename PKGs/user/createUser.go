package user

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"liteDB/PKGs/variables"
	"log"
	"os"
)


type User struct{
	Name					string						`json:"name"`
	Permission				string	 					`json:"permission"`
	Password				string						`json:"password"`
}

type PasswordManager struct {
	Name					string						`json:"name"`
	Password				string						`json:"password"`
}


/// creating a new user for the database
func CreateUser(name, password, permission string) interface{} {
	if permission != variables.Permissions[0] || permission != variables.Permissions[1] || permission != variables.Permissions[2] {
		var ManagedUsers []PasswordManager
		manageUser := PasswordManager{Name: name, Password: password}
		var Users []User
		var user User
	
		////  write users information into the password manager file which is in the root folder.
		exists := Exists(variables.PasswordManagerPath)
		if exists {
			content, err := ioutil.ReadFile(variables.PasswordManagerPath)
			HandleError(err)
			_ = json.Unmarshal(content, &ManagedUsers)
		}
			for _, user := range ManagedUsers {
				user.Name = name
				fmt.Println("User with the same name already exists")
				return nil
			}
			ManagedUsers = append(ManagedUsers, manageUser)
			content, err := json.Marshal(ManagedUsers)
			HandleError(err)
			err = ioutil.WriteFile(variables.PasswordManagerPath, content, 0644)
			HandleError(err)
	
	
		/// encrypt users password before  storing to the ordinary data file
		bytepassword := sha256.Sum256([]byte(password))
		user.Name = name
		user.Password = string(bytepassword[:])
		user.Permission = permission
	
	
		//// write users data to the ordinary users data file
		exists = Exists(variables.PasswordManagerPath)
		if exists {

			content, _ := ioutil.ReadFile(variables.UsersPath)
			HandleError(err)
			_ = json.Unmarshal(content, &Users)
		}
		Users = append(Users, user)
		content, err = json.Marshal(Users)
		HandleError(err)
		err = ioutil.WriteFile(variables.UsersPath, content, 0644)
		HandleError(err)
		return user
	}
	return "invalid permission type"
}


func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}



func HandleError(err error) {
	if err != nil && !os.IsNotExist(err){
		log.Panic(err)
	}
}