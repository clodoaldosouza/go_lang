package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		panic("Fail to  open database [" + err.Error() + "].")
	}

	sql := "INSERT INTO users(name, phone, active) VALUES('Julia Ramone', '+17985658565896585', false)"
	res, err := db.Exec(sql)
	if err != nil {
		panic("Fail to insert into [" + err.Error() + "].")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic("Fail to acquire the lastId [" + err.Error() + "].")
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
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

*/
