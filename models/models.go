package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_DB_USER       = "root"
	_DB_PASS       = "root"
	_DB_DBNAME     = "test"
	_DB_Driver     = "mysql"
	_DB_DEFAULT_DB = "default"
	//_DB_PROTOCOL   = ""
	//_DB_ADDR       = "localhost"
	_DB_CHARSET = "utf8"
)

type Categroy struct {
	Id       int
	Title    string    `orm:"unique"`
	Dateline time.Time `orm:"index"`
}

func RegisterDB() {
	orm.RegisterModel(new(Categroy))
	orm.RegisterDriver(_DB_Driver, orm.DR_MySQL)
	orm.RegisterDataBase(_DB_DEFAULT_DB, _DB_Driver, _DB_USER+":"+_DB_PASS+"@"+"/"+_DB_DBNAME+"?"+"charset="+_DB_CHARSET)
}
