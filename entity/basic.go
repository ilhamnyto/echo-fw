package entity

type Paging struct {
	Next	bool		`json:"next"`
	Cursor	string		`json:"cursor"`
}

type DataResponse struct {
	Data   interface{} `json:"data"`
    Paging interface{} `json:"paging,omitempty"`
}