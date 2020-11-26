package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB



func initDB() (err error) {
	dsn := "root:lzh123456@tcp(39.102.112.15:3306)/bms?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryOne(sql string,args ...interface{})(b Book,err error) {

	err = db.Get(&b, sql, args...)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	//fmt.Printf("id:%d name:%s price:%d\n", b.Id,b.BookName,b.Price)
	fmt.Printf("%v\n",b)
	return b,err
}

// 查询多条数据示例
func queryMulti(sql string,args ...interface{})(bookList []*Book,err error) {
	err = db.Select(&bookList, sql, args...)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("books:%#v\n", )
	return
}

//增删改
func modify(sql string, args ...interface{}) (err error) {

	ret, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Printf("modify failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID or rows affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("modify success, the id is %d.\n", theID)
	return

}

