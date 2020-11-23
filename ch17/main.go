package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type User struct {
	Id       int
	UserName string
	Password string
}

const (
	USERNAME = "root"
	PASSWORD = "12345678"
	IP       = "127.0.0.1"
	PORT     = "3306"
	dbName   = "loginserver"
)

var DB *sql.DB

func main() {
	InitDB()
	beijing := User{
		UserName: "beijing",
		Password: "beijing",
	}
	shanghai := User{
		UserName: "shanghai",
		Password: "shanghai",
	}
	InsertUser(beijing)
	InsertUser(shanghai)

	for _, user := range SelectAllUser() {
		fmt.Println(user.UserName + " " + user.Password)
	}
}

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}

func InsertUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO nk_user (`name`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.UserName, user.Password)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM nk_user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func UpdateUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE nk_user SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.UserName, user.Password, user.Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func SelectUserById(id int) User {
	var user User
	err := DB.QueryRow("SELECT id,name,password FROM nk_user WHERE id = ?", id).Scan(&user.Id, &user.UserName, &user.Password)
	if err != nil {
		fmt.Println("查询出错了", err.Error())
	}
	return user
}

func SelectAllUser() []User {
	//执行查询语句
	rows, err := DB.Query("SELECT id,name,password from nk_user")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var users []User
	//循环读取结果
	for rows.Next() {
		var user User
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}
