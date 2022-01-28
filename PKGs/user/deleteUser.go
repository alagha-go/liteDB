package user

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io/ioutil"
	"liteDB/PKGs/variables"
)


/// func to delete a user
func DeleteUser(name, password string) error {
	var oldUsers []User
	var newUsers []User
	verified, err := VerifyUser(name, password, "read")
	if !verified{
		return err
	}

	content, err := ioutil.ReadFile(variables.UsersPath)
	HandleError(err)
	err = json.Unmarshal(content, &oldUsers)
	HandleError(err)
	
	for _, user := range oldUsers {
		if user.Name != name {
			newUsers = append(newUsers, user)
		}
	}
	content, err = json.Marshal(newUsers)
	HandleError(err)
	err = ioutil.WriteFile(variables.UsersPath, content, 0644)
	HandleError(err)
	
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

	bytes := sha256.Sum256([]byte(password))
	hashedPassword := string(bytes[:])

	for _, user := range AllUsers {
		if user.Name == name {
			exist = true
			if user.Password == hashedPassword{
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