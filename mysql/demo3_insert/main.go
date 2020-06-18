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

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connect to mysql failed:", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //空闲连接的最大存活时间
	DB.SetMaxOpenConns(100)                  //最大连接数
	CreateTable(DB)
	InsertData(DB)
}

func InsertData(DB *sql.DB) {
	sql := `INSERT INTO student(name,age) values (?,?)`
	result, err := DB.Exec(sql, "resolmi", "21")
	if err != nil {
		fmt.Println("insert data failed:", err)
		return
	}
	lastInsertId, err := result.LastInsertId() //获取插入数据的自增id
	if err != nil {
		fmt.Println("get lastInsertId (auto incr) failed")
	}
	fmt.Println("lastInsertId is：", lastInsertId)
	rowAffect, err := result.RowsAffected() //获取受影响的行数
	if err != nil {
		fmt.Println("get rowAffect failed:", err)
	}
	fmt.Println("rowAffect is：", rowAffect)
}

func CreateTable(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS student (
			  id int(11) NOT NULL AUTO_INCREMENT,
			  name varchar(255) DEFAULT NULL,
			  age int(11) DEFAULT NULL,
			  PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table succeeded!")
}
