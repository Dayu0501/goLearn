package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

func init() {
	OpenDB()
}

func OpenDB() {
	database, err := sqlx.Open("mysql", "root:em-data-9527@tcp(192.168.70.210:6033)/chejian_refactor_v6")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}

func DbInsert(sql string) error {
	r, err := Db.Exec(sql)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err
	}

	fmt.Println("insert succ:", id)
	return nil
}

func DbSelect() {
	OpenDB()

	var person []user
	err := Db.Select(&person, "select * from tab_no_index")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}

func QueryDb(sql string) ([]string, error) {

	var setPic []string
	err := Db.Select(&setPic, sql)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}

	fmt.Println("select succ:", setPic)
	return setPic, nil
}