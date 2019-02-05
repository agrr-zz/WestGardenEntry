package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//表的设计
type User struct {
	Id      int
	Name    string `orm:"unique"`
	Pwd     string
	Article []*Article `orm:"rel(m2m)"`
}

//文章结构体
type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:"default(0);null"`
	Acontent string
	Aimg     string
	Atype    string

	ArticleType *ArticleType `orm:"rel(fk)"`
	User        []*User      `orm:"reverse(many)"`
}

//类型表
type ArticleType struct {
	Id      int
	Tname   string
	Article []*Article `orm:"reverse(many)"`
}

type EventList struct {
	Id             int
	Eventname      string
	EventStartDate string
	EventStartTime string
	EventContinue  string
	// InDate         time.Time `orm:"auto_now"`
}

type MemberList struct {
	Id         int
	MemberName string
	Phone      string
}

func init() {
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User), new(Article), new(ArticleType), new(EventList), new(MemberList))
	// 生成表
	orm.RunSyncdb("default", false, true)
}
