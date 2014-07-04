package models

import (
	"fmt"
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
	Cid             int64
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

type Reply struct {
	Id        int64
	Content   string `orm:"size(1000)"`
	Tid       int64
	Title     string
	NickName  string
	ReplyTime time.Time
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic), new(Reply))
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

func AddTopic(title, content, category string) error {

	cid, err := strconv.ParseInt(category, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()

	Topic := &Topic{Title: title, Content: content, Cid: cid, Created: time.Now(), Author: "admin"}

	_, err = o.Insert(Topic)

	if err != nil {
		return err
	}

	return nil
}

func GetAllTopic() ([]orm.Params, error) {

	o := orm.NewOrm()
	var maps []orm.Params

	num, err := o.Raw("SELECT topic.*, category.title AS CateTitle FROM topic LEFT JOIN category ON (topic.cid = category.id)").Values(&maps)
	fmt.Println(maps)
	if err == nil && num > 0 {
		return maps, err
	}
	return maps, err
}

func GetOneTopic(id string) (*Topic, error) {

	cid, _ := strconv.ParseInt(id, 10, 64)
	topic := &Topic{Id: cid}

	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err := qs.Filter("id", cid).One(topic)

	return topic, err
}

func UpdateTopic(id, title, content, category string) error {

	o := orm.NewOrm()
	num, err := o.QueryTable("topic").Filter("id", id).Update(orm.Params{"title": title, "content": content, "cid": category})

	if num > 0 {
		return nil
	}

	return err
}

func GetOneTopicAndCategory(id string) ([]orm.Params, error) {

	o := orm.NewOrm()
	var res []orm.Params

	_, err := o.Raw("SELECT topic.*, category.title AS CateTitle FROM topic LEFT JOIN category ON (topic.cid = category.id) WHERE topic.id = ?", id).Values(&res)

	return res, err
}

func AddReply(title, content, nickname, tid string) error {

	intTid, err := strconv.ParseInt(tid, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()

	reply := &Reply{Title: title, Content: content, NickName: nickname, Tid: intTid, ReplyTime: time.Now()}

	_, err = o.Insert(reply)

	if err != nil {
		return err
	}

	return nil
}

func GetTopicReplyByTopicId(tid string) ([]orm.Params, error) {

	o := orm.NewOrm()
	var res []orm.Params

	_, err := o.Raw("SELECT reply.* FROM reply WHERE reply.tid = ?", tid).Values(&res)

	return res, err
}

func GetTopicByCategoryId(cid, id interface{}) ([]orm.Params, error) {

	o := orm.NewOrm()
	var res []orm.Params

	var err error
	if id == "" {
		_, err = o.Raw("SELECT topic.* FROM topic WHERE topic.cid = ?", cid).Values(&res)
	} else {
		_, err = o.Raw("SELECT topic.* FROM topic WHERE topic.cid = ? AND id <> ?", cid, id).Values(&res)
	}

	return res, err
}
