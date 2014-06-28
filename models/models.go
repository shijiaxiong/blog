package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

type Category struct {
	Id         int64
	Title      string    `orm:"unique"`
	Dateline   time.Time `orm:"index"`
	Views      int64     `orm:"index"`
	TopicTime  time.Time `orm:"index"`
	TopicCount int64     `orm:"index"`
	LastUserId int64     `orm:"index"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_DB_Driver, orm.DR_MySQL)
	orm.RegisterDataBase(_DB_DEFAULT_DB, _DB_Driver, _DB_USER+":"+_DB_PASS+"@"+"/"+_DB_DBNAME+"?"+"charset="+_DB_CHARSET)
}

func AddCategory(title string) error {

	o := orm.NewOrm()

	cate := &Category{Title: title, Dateline: time.Now()}

	qs := o.QueryTable("category")

	err := qs.Filter("title", title).One(cate)

	if err == nil {

		return err
	}

	_, err = o.Insert(cate)

	if err != nil {

		return err
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {

	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}

	_, err = o.Delete(cate)

	return err
}
