package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"userid"`
	Username string `db:"ssername"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:cgs123456@tcp(127.0.0.1:3306)/studb")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu004", 2)
	if err != nil {
		fmt.Println("Exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
	}
	fmt.Println("update succ:", row)

	//	Delete 操作
	res1, err := Db.Exec("delete  from person where user_id=?", 2)
	if err != nil {
		fmt.Println("Exec failed,", err)
		return
	}
	row1, err := res1.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
	}
	fmt.Println("update succ:", row1)
}
