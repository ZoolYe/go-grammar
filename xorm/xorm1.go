package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"
)

var engine *xorm.Engine

type userinfo struct {
	Uid        int       //`xorm:"not null pk autoincr INT(11)"`
	Username   string    //`xorm:"VARCHAR(32)"`
	Department string    //`xorm:"VARCHAR(32)"`
	Created    time.Time //`xorm:"DATE"`
}

func init() {

	var err error
	//创建orm引擎
	engine, err = xorm.NewEngine("mysql", "root:123456@/go?charset=utf8")
	if err != nil {
		fmt.Println("创建引擎失败")
		return
	}

	if err := engine.Ping(); err != nil {
		fmt.Println("测试连接失败")
		return
	}

	//打印SQL日志
	engine.ShowSQL()
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})
}

func main() {
	//insert(userinfo{Username: "sixi", Department: "技术部", Created: time.Now()})
	//delete(userinfo{Username:"sixi"})
	//update(userinfo{Username: "sixi", Department: "大中台"})

	user := new(userinfo)
	user.Uid = 16401
	getUserInfo(user)
	fmt.Println(user.Uid, user.Username, user.Department, user.Created)
}

func insert(u userinfo) {
	engine.Insert(u)
	fmt.Println(time.Now())
}

func delete(u userinfo) {
	engine.Delete(u)
}

func update(u userinfo) {
	engine.Where("username=?", u.Username).Update(u)
}

func getUserInfo(u *userinfo) {
	engine.Get(u)
}
