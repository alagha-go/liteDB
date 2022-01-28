package main

import (
	"fmt"
	"net/http"

	"liteDB/PKGs/user"
	"liteDB/PKGs/variables"
)


func main(){
	variables.Starter()
	fmt.Println("starting litedb...")

	user.CreateUser("Abdihakim", "Alagha", "sudo")
	user.ReadAllUsers()

	http.HandleFunc("/", Hello)


	http.ListenAndServe(fmt.Sprintf("%s%s", variables.IP(), variables.PORT()), nil)
}

func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>It looks like you are trying to access litedb over the http Network</h1>"))
}