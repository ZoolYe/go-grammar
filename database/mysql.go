package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var db *sql.DB

func init() {
	var e error
	//打开数据库链接
	db, e = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go?charset=utf8&parseTime=true&loc=Local")
	checkErr(e)
}

func main() {

	//insert()
	//update()
	selectSql()
	delete()
}

func insert() {
	//准备sql
	stmt, e := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(e)
	//执行准备好的sql语句
	exec, e := stmt.Exec("zoolye", "技术部", time.Now().UTC())
	checkErr(e)
	//获取插入的主键id
	lastInsertId, e := exec.LastInsertId()
	checkErr(e)
	fmt.Println("新增数据的uid是：" + strconv.FormatInt(lastInsertId, 10))
	//获取影响的行数
	rowsAffected, e := exec.RowsAffected()
	checkErr(e)
	fmt.Println("新增了" + strconv.FormatInt(rowsAffected, 10) + "条数据")
}

func update() {

	stmt, e := db.Prepare("UPDATE userinfo SET department = ? WHERE uid = ?")
	checkErr(e)

	exec, e := stmt.Exec("ERP开发小组", 16394)
	checkErr(e)
	rowsAffected, e := exec.RowsAffected()
	fmt.Println("成功更新" + strconv.FormatInt(rowsAffected, 10) + "条数据")

}

func selectSql() {

	rows, e := db.Query("SELECT * FROM userinfo")
	checkErr(e)

	for rows.Next() {
		userInfo := new(userInfo)
		e := rows.Scan(&userInfo.uid, &userInfo.username, &userInfo.department, &userInfo.created)
		checkErr(e)
		fmt.Println(userInfo)
	}

}

func delete() {
	stmt, e := db.Prepare("DELETE FROM userinfo WHERE uid = ?")
	checkErr(e)
	exec, e := stmt.Exec("16394")
	checkErr(e)
	rowsAffected, e := exec.RowsAffected()
	fmt.Println("删除了" + strconv.FormatInt(rowsAffected, 10) + "条数据")

}

type userInfo struct {
	uid        int
	username   string
	department string
	created    string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
