package models

type Responses struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}

type Pagination struct {
	Current  int `json:"current_page"`
	Next     int `json:"next_page"`
	Pagesize int `json:"page_size"`
}
