package variables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	ManagerPath = "/data/root/liteDB/.config/"
	PasswordManagerPath = "/data/root/liteDB/.config/users.json"
	UsersPath = "/data/db/liteDB/.users/users.json"
	Permissions []string = []string{"sudo", "read", "write"}
)



type Config struct {
	IPAddress 						string						`json:"ip_address"`
	PORT 							string						`json:"port"`
	DBPath							string						`json:"db_path"`
}



/// func to get path where the data should be stored
func DBPath() string {
	var config Config
	content, err := ioutil.ReadFile(fmt.Sprintf("%sconfig.json", ManagerPath))
	HandleError(err)
	err = json.Unmarshal(content, &config)
	HandleError(err)

	return config.DBPath
}


/// func to get the ip address to run the application
func IP() string {
	var config Config
	content, err := ioutil.ReadFile(fmt.Sprintf("%sconfig.json", ManagerPath))
	HandleError(err)
	err = json.Unmarshal(content, &config)
	HandleError(err)

	if config.IPAddress == "" {
		return "127.0.0.1"
	}

	return config.IPAddress
}


/// func to get port to run the application
func PORT() string {
	var config Config
	content, err := ioutil.ReadFile(fmt.Sprintf("%sconfig.json", ManagerPath))
	HandleError(err)
	err = json.Unmarshal(content, &config)
	HandleError(err)

	if config.PORT == "" {
		return ":1480"
	}

	return config.PORT
}


func Starter() {
	exists := Exists(ManagerPath)
	if !exists {
		err := os.MkdirAll("/data/root/liteDB/.config", 0644)
		HandleError(err)
		err = os.MkdirAll("/data/db/liteDB/.users", 0644)
		HandleError(err)
	}
	exists = Exists(fmt.Sprintf("%sconfig.json", ManagerPath))
	if !exists {
		config := Config{IPAddress: "127.0.0.1", PORT: ":1480", DBPath: "/data/db/liteDB/"}
		content, err := json.Marshal(config)
		HandleError(err)
		err = ioutil.WriteFile(fmt.Sprintf("%sconfig.json", ManagerPath), content, 0644)
		HandleError(err)
	}
	exists = Exists(DBPath())
	if !exists {
		err := os.MkdirAll(DBPath(), 0644)
		HandleError(err)
	}
}

/// func to check if a directory or file exists
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}


/// func to handle all errors in this pkg
func HandleError(err error) {
	if err != nil && !os.IsExist(err){
		log.Panic(err)
	}
}