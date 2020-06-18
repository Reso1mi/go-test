package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type User struct {
	//gorm 后添加约束
	Id   int    `gorm: "size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Age  int    `gorm: "size:11;DEFAULT NULL" json:"age"`
	Name string `gorm: "size:255;DEFAULT NULL" json:"name"`
}

func main() {
	var (
		MysqlDB *gorm.DB
		err     error
	)
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	if MysqlDB, err = gorm.Open("mysql", conn); err != nil {
		panic("数据库连接失败")
	}
	defer MysqlDB.Close()
	//建表的时候全局禁用表名复数，默认struct User的表名是users，如果通过TableName设置的就没有影响
	MysqlDB.SingularTable(true)
	//建表
	MysqlDB.AutoMigrate(&User{})
	fmt.Println("create table success")
	fmt.Println("---------------")
	AddUser(MysqlDB)
	fmt.Println("---------------")
	UpdateUser(MysqlDB)
	fmt.Println("---------------")
	DeleteUser(MysqlDB)
}

func AddUser(DB *gorm.DB) {
	user := &User{
		Name: "imlgw",
		Age:  21,
	}
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println("add new User failed: ", err)
	}
	fmt.Println("add user success！！！")
	GetUser(DB, user.Id)
}

func UpdateUser(DB *gorm.DB) {
	user := &User{
		Name: "imlgw",
		Age:  2111111,
	}
	if err := DB.Save(&user).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println("Update success！！！")
	GetUser(DB, user.Id)
}

func GetUser(DB *gorm.DB, id int) {
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		fmt.Println("Query user failed:", err)
	}
	fmt.Printf("Get User [id = %d: %v]", id, user)
}

func DeleteUser(DB *gorm.DB) {
	user := &User{Id: 2}
	if err := DB.Delete(&user).Error; err != nil {
		fmt.Println("delete user failed:", err)
	}
	fmt.Println("delete user success")
}
