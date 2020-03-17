package services

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

var baseURL string = beego.AppConfig.String("elasticsearch_url")

// PayloadLogES is object payload
type PayloadLogES struct {
	ID      uint      `json:"id"`
	Created time.Time `json:"created"`
}

// ResponseMapping is a response for mapping
type ResponseMapping struct {
	Acknowledged bool `json:"acknowledged"`
}

// ES is service elasticsearch
type ES struct {
	Id      uint
	Created time.Time
}

// Mapping is function for create index elasticsearch
func (o ES) Mapping() ResponseMapping {
	var rsMapping ResponseMapping

	url := fmt.Sprintf("%s/news", baseURL)
	req := httplib.Put(url)
	req.Header("Content-Type", "application/json")

	req.ToJSON(&rsMapping)

	return rsMapping
}

// CreateLogNews is function for create news log
func (o ES) CreateLogNews() interface{} {
	var payload PayloadLogES
	var rs interface{}

	mapping := o.Mapping()

	if mapping.Acknowledged {
		fmt.Println("already mapping")
	}

	url := fmt.Sprintf("%s/news/_doc?refresh", baseURL)

	payload.Created = o.Created
	payload.ID = o.Id

	req := httplib.Post(url)
	req.JSONBody(payload)

	req.Header("Content-Type", "application/json")

	req.ToJSON(&rs)
	return rs
}
