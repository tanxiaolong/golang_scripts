package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type DBWorker struct {
	Dsn string
}

func Conn() *sql.DB {
	dbw := DBWorker{
		Dsn: "root:letv@tcp(10.110.80.71)/test",
	}
	db, err := sql.Open("mysql", dbw.Dsn)
	//defer db.Close()
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		panic(err)
		return nil
	}
	return db
}

func main() {
	db := Conn()
	fmt.Printf("%+v\n", db)
	sql := "select * from coupon"
	rlt, err := db.Query(sql)
	fmt.Printf("%+v", rlt, err)
	for rlt.Next() {
		var (
			id     string
			cpcode string
		)
		if err := rlt.Scan(&id, &cpcode); err != nil {
			fmt.Println(err)
		}
		fmt.Println(id, cpcode)
	}
	defer rlt.Close()
}
