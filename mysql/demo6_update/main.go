package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "admin"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "test"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connect to mysql failed:", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //空闲连接的最大存活时间
	DB.SetMaxOpenConns(100)                  //最大连接数
	UpdateData(DB)
	QueryMany(DB)
}

func UpdateData(DB *sql.DB) {
	sql := `UPDATE student set name=? where id=?`
	if _, err := DB.Exec(sql, "imlgw", 2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("update data succeeded!!!")
}

func QueryMany(DB *sql.DB) {
	sql := `SELECT id,name,age from student`
	stu := &Student{}
	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		err = rows.Scan(&stu.Id, &stu.Name, &stu.Age) //一行行的扫描
		if err != nil {
			fmt.Println("scan failed: ", err)
			return
		}
		fmt.Println("scan success!!!", stu)
	}
}
