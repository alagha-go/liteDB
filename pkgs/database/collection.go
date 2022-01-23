package database

import (
	"fmt"
	"os"
)

type Database struct{
	Name 						string
	Collections					[]Collection
}

type DB struct {
	Database					string
	Collection					string
}

type Collection struct {
	Name 						string
	Documents					[]string
}

const (
	DBPath = "./DBs/"
)


func CreateDatabase(name string) string {
	database := fmt.Sprintf("%s%s", DBPath, name)
	exists := DBexists(database)
	if !exists {
		err := os.Mkdir(database, 0755)
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Printf("database %s successfully created\n", name)
		}
	}else{
		fmt.Printf("database %s already exists\n", name)
	}
	return name
}


func DBexists(DBFile string) bool {
	if _, err := os.Stat(DBFile); os.IsNotExist(err){
		return false
	}
	return true
}