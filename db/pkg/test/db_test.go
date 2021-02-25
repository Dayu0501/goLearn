package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

type Person1 struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

type user struct {
	Id 		int
	Name 	string
}

var Db *sqlx.DB

func OpenDB() {
	database, err := sqlx.Open("mysql", "root:em-data-9527@tcp(192.168.20.71:3306)/db_test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

func TestDbInsert(t *testing.T) {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func TestDbSelect(t *testing.T) {
	OpenDB()

	var person []user
	err := Db.Select(&person, "select * from tab_no_index")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}