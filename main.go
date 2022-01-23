package main

import (
	"encoding/json"
	"fmt"
	"liteDB/pkgs/database"
)


type User struct {
	ID							string						`json:"_id,omitempty"`
	FirstName					string						`json:"first_name,omitempty"`
	LastName					string						`json:"last_name,omitempty"`
	Age							int							`json:"age,omitempty"`
}



func main(){
	var users []User
	// var user User
	// user = User{"", "Another New", "User", 9}
	db := database.CreateDatabase("Interphlix").CreateCollection("Users")
	// fmt.Println(user)
	// cursor := db.CreateDocument(user)
	// cursor.Decode(&user)
	// fmt.Println(user)
	db.GetAllDocuments().DecodeMany(&users)
	fmt.Println(len(users))
	bytedUsers, _ := json.Marshal(users)
	fmt.Println(string(bytedUsers))
}