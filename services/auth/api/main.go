package main

import (
	"github.com/admgo/admgo/services/auth/db"
)

func main() {

	err := db.Init()
	if err != nil {
		panic("can not init database")
	}
	dbw := db.DBW()
	type Result struct {
		Host string
		user string
	}

	var result Result
	dbw.Raw("SELECT host, user FROM user WHERE user = ?", "addmgo").Scan(&result)
	//println(result)
}
