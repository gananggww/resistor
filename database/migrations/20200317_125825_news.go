package main

import (
	"github.com/astaxie/beego/migration"
)

//News_20200317_125825 is DO NOT MODIFY
type News_20200317_125825 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &News_20200317_125825{}
	m.Created = "20200317_125825"

	migration.Register("News_20200317_125825", m)
}

// Up is Run the migrations
func (m *News_20200317_125825) Up() {
	// m.CreateTable("news", "InnoDB", "utf8")
	// m.PriCol("id").SetAuto(true).SetNullable(false).SetDataType("INT").SetUnsigned(true)
	// m.PriCol("author").SetDataType("TEXT")
	// m.PriCol("body").SetDataType("TEXT")
	m.SQL("CREATE TABLE news (id serial PRIMARY KEY, author TEXT NOT NULL, body TEXT NOT NULL, created TIMESTAMP NOT NULL default current_timestamp)")
}

//Down is Reverse the migrations
func (m *News_20200317_125825) Down() {
	m.SQL("DROP TABLE news")
}
