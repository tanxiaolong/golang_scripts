package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "sync"
	_ "time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:Qq@261337699@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("err: %+v\n", err)
		panic("连接数据库失败")
	}
	defer db.Close()

	//wg := sync.WaitGroup{}
	// 创建
	sql := `insert into products(code, price) values(?,?) on duplicate key update id=?,code=?,price=?`
	for i := 0; i < 4864; i++ {
		//wg.Add(1)
		//go func(i int) {
		//defer wg.Done()
		db.Exec(sql, fmt.Sprintf("L%d", i), i, i, fmt.Sprintf("L%d", i), i)
		//}(i)
	}
	//wg.Wait()
}
