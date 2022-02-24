package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	usermodel "souza.com/index/mysql.example/models"
)

func main() {

	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		panic("Fail to  open database [" + err.Error() + "].")
	}

	helenId := 3
	res, err := db.Query("SELECT u.* FROM users as u WHERE id = ?", helenId)
	defer res.Close()
	if err != nil {
		panic("Fail to select [" + err.Error() + "].")
	}

	if res.Next() {
		user := usermodel.User{}
		err := res.Scan(&user.Id, &user.Name, &user.Phone, &user.Active)
		if err != nil {
			panic("Fail [" + err.Error() + "].")
		}

		fmt.Printf("%v\n", user)
	} else {

		fmt.Println("No user found")
	}
}
