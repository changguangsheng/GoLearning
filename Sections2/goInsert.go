package main

import "fmt"
import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	EMail    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	Citry   string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:cgs123456@tcp(127.0.0.1:3306)/studb")
	if err != nil {
		fmt.Println("open mysql filed, ", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu003", "girl", "stu03@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("inser sycc: ", id)
}
