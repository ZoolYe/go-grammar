package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ormer orm.Ormer

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql",
		"root:123456@tcp(localhost:3306)/go?charset=utf8&parseTime=true&loc=Local", 30)

	//注册定义的model
	//也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))
	orm.RegisterModel(new(User))
	// 创建table
	orm.RunSyncdb("default", false, true)
	//开启打印调试
	orm.Debug = true
	//打开数据库链接
	ormer = orm.NewOrm()
}

type User struct {
	Id      int `orm:"column(id);pk"` // 设置主键
	Name    string
	Sex     string
	IdCard  string
	Address string
	Wechat  string
	Phone   string
}

func main() {
	//insert()
	//delete()
	//update()
	//get()
	//insertMulti()
	//update2()
}

func insert() {

	user := User{Name: "zoolye", Sex: "男", IdCard: "421124199906132032",
		Address: "浙江杭州", Wechat: "phonelove", Phone: "17605881104"}

	insert, e := ormer.Insert(&user)
	fmt.Println(insert, e)
}

func insertMulti() {

	users := []User{
		{Name: "zoolye", Sex: "男"},
		{Name: "zoolye", Sex: "男"},
		{Name: "zoolye", Sex: "男"},
	}
	//第一个参数 bulk 为并列插入的数量
	id, e := ormer.InsertMulti(100, users)
	fmt.Println(id, e)
}

func delete() {
	user := User{Id: 6}
	ormer.Delete(&user)
}

func update() {
	user := User{Id: 7, Sex: "女", Name: "陈濛"}
	//默认更新全部字段
	//ormer.Update(&user)
	//只更新制定的字段
	ormer.Update(&user, "Sex", "Name")
}

func update2() {
	user := User{Id: 7}
	err := ormer.Read(&user)
	if err == nil {
		user.Name = "叶钊"
		ormer.Update(&user)
	}
}

func get() {
	user := User{Id: 7}
	err := ormer.Read(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
