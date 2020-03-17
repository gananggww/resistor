package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

// News is a model  object for  news
type News struct {
	Id      uint      `orm:"pk;auto"`
	Author  string    `orm:"size(255)"`
	Body    string    `orm:"size(255)"`
	Created time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(News))
}

// GetNews is a function for get all a news
func GetNewsCount() (v []*News, err error) {
	o := orm.NewOrm()

	var news []*News
	_, errorGet := o.QueryTable("news").All(&news)

	return news, errorGet
}

// GetNews is a function for get all pagination a news
func GetNews(limit int, offset int) (v []*News, err error) {
	o := orm.NewOrm()

	var news []*News
	qs, errorGet := o.QueryTable("news").Limit(limit).Offset(offset).All(&news)

	fmt.Println(qs, news)

	return news, errorGet

}

// CreateNews is a function for create a news
func CreateNews(news News) (v *News, err error) {
	o := orm.NewOrm()

	news.Created = time.Now()

	_, errorCreate := o.Insert(&news)

	if errorCreate != nil {
		fmt.Println("==ERROR==", errorCreate)
	}
	return &news, errorCreate
}
