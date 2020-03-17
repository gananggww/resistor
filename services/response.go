package services

// Response is response a api
type Response struct {
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type Paginate struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
	Count   int `json:"count"`
}

type ResponseP struct {
	Msg      string      `json:"msg"`
	Status   bool        `json:"status"`
	Data     interface{} `json:"data"`
	Paginate Paginate    `json:"paginate"`
}
