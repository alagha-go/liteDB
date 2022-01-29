package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"liteDB/PKGs/variables"
)

/// func to delete a user
func DeleteUser(name, password string) error {
	var oldUsers []User
	var oldManagedUsers []PasswordManager
	var newUsers []User
	var newManagedUsers []PasswordManager
	verified, err := VerifyUser(name, password, "read")
	if !verified{
		return err
	}

	content, err := ioutil.ReadFile(variables.UsersPath)
	content1, err1 := ioutil.ReadFile(variables.PasswordManagerPath)
	HandleError(err)
	HandleError(err1)
	err1 = json.Unmarshal(content1, &oldManagedUsers)
	HandleError(err1)
	err = json.Unmarshal(content, &oldUsers)
	HandleError(err)
	
	for _, user := range oldUsers {
		if user.Name != name {
			newUsers = append(newUsers, user)
		}
	}


	for _, user := range oldManagedUsers {
		if user.Name != name {
			newManagedUsers = append(newManagedUsers, user)
		}
	}

	content, err = json.Marshal(newUsers)
	HandleError(err)
	content1, err1 = json.Marshal(newManagedUsers)
	HandleError(err1)
	err = ioutil.WriteFile(variables.UsersPath, content, 0644)
	HandleError(err)
	err1 = ioutil.WriteFile(variables.PasswordManagerPath, content1, 0644)
	HandleError(err1)
	
	return nil
}



//// func to verify if user has the right permissions.
func VerifyUser(name, password, permission string) (bool, error) {
	var exist bool
	var noPermission bool
	var AllUsers []User
	content, err := ioutil.ReadFile(variables.UsersPath)
	HandleError(err)
	err = json.Unmarshal(content, &AllUsers)
	HandleError(err)

	

	for _, user := range AllUsers {
		if user.Name == name {
			exist = true
			if CompareHash(password, user.Password){
				if user.Permission == "sudo" {
					return true, nil
				}else if user.Permission == permission {
					return true, nil
				}else if user.Permission == "write" && permission == "read" {
					return true, nil
				}else {
					noPermission = true
				}
			}
		}
	}

	if noPermission{
		return false, errors.New("user does not have the right permissions")
	}

	if exist{
		return false, errors.New("wrong password")
	}

	return false, errors.New("user does not exist")
}