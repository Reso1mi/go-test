package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/*
	为了测试批量插入的新功能，采用了新的v2版本
*/

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
	Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Age  int    `gorm:"size:11;DEFAULT NULL" json:"age"`
	Name string `gorm:"size:255;DEFAULT NULL" json:"name"`
}

func main() {
	var (
		MysqlDB *gorm.DB
		err     error
	)
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	if MysqlDB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		//取代之前的 MysqlDB.SingularTable(true)
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}); err != nil {
		panic("数据库连接失败")
	}
	//defer MysqlDB.Close() v2这个没有了，文档里面也没写close的部分
	//建表的时候全局禁用表名复数，默认struct User的表名是users，如果通过TableName设置的就没有影响
	//MysqlDB.SingularTable(true) 这个也没有了,取而代之的是在上面的Config中设置 这里还踩了个坑，后面直接建了一个users的表没注意，一直在看user表
	//建表
	MysqlDB.AutoMigrate(&User{})
	fmt.Println("create table success")
	fmt.Println("---------------")
	AddUser(MysqlDB)
	fmt.Println("---------------")
	UpdateUser(MysqlDB)
	fmt.Println("---------------")
	DeleteUser(MysqlDB)
	fmt.Println("---------------")
	AddManyUser(MysqlDB)
	fmt.Println("---------------")
	AddManyUser2(MysqlDB)
}

func AddUser(DB *gorm.DB) {
	user := &User{
		Name: "imlgw",
		Age:  21,
	}
	if err := DB.Create(&user).Error; err != nil {
		fmt.Println("add new User failed: ", err)
		return
	}
	fmt.Println("add user success！！！")
	GetUser(DB, user.Id)
}

//批量插入测试
func AddManyUser(DB *gorm.DB) {
	users := []User{{Name: "imlgw1", Age: 11}, {Name: "imlgw2", Age: 12}, {Name: "imlgw3", Age: 13}}
	if err := DB.Create(users).Error; err != nil {
		fmt.Println("add  users failed: ", err)
		return
	}
	fmt.Println("add  users success！！！")
	for _, user := range users {
		GetUser(DB, user.Id)
	}
}

//这种是不行的，必须知道具体类型
func AddManyUser2(DB *gorm.DB) {
	user1 := User{Name: "1resolmi", Age: 11}
	user2 := User{Name: "2resolmi", Age: 12}
	user3 := User{Name: "3resolmi", Age: 12}
	var users []interface{}
	users = append(users, user1, user2, user3)
	if err := DB.Create(users).Error; err != nil {
		fmt.Println("add  users 2 failed: ", err)
		return
	}
	fmt.Println("add  users 2 success！！！")
	for _, user := range users {
		GetUser(DB, user.(User).Id)
	}
}

func UpdateUser(DB *gorm.DB) {
	user := &User{
		Name: "imlgw",
		Age:  2111111,
	}
	if err := DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update success！！！")
	GetUser(DB, user.Id)
}

func GetUser(DB *gorm.DB, id int) {
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		fmt.Println("Query user failed:", err)
		return
	}
	fmt.Printf("Get User [id = %d: %v]", id, user)
}

func DeleteUser(DB *gorm.DB) {
	user := &User{Id: 2}
	if err := DB.Delete(&user).Error; err != nil {
		fmt.Println("delete user failed:", err)
		return
	}
	fmt.Println("delete user success")
}
