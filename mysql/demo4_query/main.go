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
	//表已经建好
	//QueryOne(DB)
	QueryMany(DB)
}

func QueryOne(DB *sql.DB) {
	//注意这里返回的顺序
	sql := `SELECT name,id,age from student where id=?`
	stu := &Student{}
	row := DB.QueryRow(sql, 1)
	//注意顺序和上面sql返回的顺序要一致,注意传指针,这样才能注入值,看了一圈源码感觉良好
	if err := row.Scan(&stu.Name, &stu.Id, &stu.Age); err != nil {
		fmt.Println("scan failed: ", err)
		return
	}
	fmt.Printf("id=1 row data is: %v", stu)
}

func QueryMany(DB *sql.DB) {
	sql := `SELECT id,name,age from student where name=?`
	stu := &Student{}
	rows, err := DB.Query(sql, "resolmi")
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
