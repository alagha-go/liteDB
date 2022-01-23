package main

import (
	"liteDB/pkgs/database"
)



func main(){
	database.GenerateId()
	database.CreateDatabase("Alagha")
}