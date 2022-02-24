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

	res, err := db.Query("SELECT u.* FROM users as u ORDER BY u.id ")
	defer res.Close()
	if err != nil {
		panic("Fail to select [" + err.Error() + "].")
	}

	for res.Next() {
		user := usermodel.User{}
		err := res.Scan(&user.Id, &user.Name, &user.Phone, &user.Active)
		if err != nil {
			panic("Fail [" + err.Error() + "].")
		}

		fmt.Printf("%v\n", user)
	}
}

// Download lib  : go get -u github.com/go-sql-driver/mysql
// Simple docker : docker run --name mymysql -p 3306:3306 -e MYSQL_USER:user -e MYSQL_PASSWORD:secret -e MYSQL_ROOT_PASSWORD=toor -d mysql:latest

/** SQL Command
DROP DATABASE IF EXISTS test;
CREATE DATABASE test;
USE test;
DROP TABLE IF EXISTS users;
CREATE TABLE users(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), phone VARCHAR(30), active BOOLEAN);
INSERT INTO users(name, phone, active) VALUES('Loren Uptown Fisher', "+1363333671000", false),('John Lipxten So Jr', "+475525012025520", true),('Helen Matthew', "+98325622522523", true);
*/

/*
Output:
{1 Loren Uptown Fisher +1363333671000 false}
{2 John Lipxten So Jr +475525012025520 true}
{3 Helen Matthew +98325622522523 true}
*/
