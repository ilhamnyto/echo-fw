package entity

type Paging struct {
	Next	bool	`json:"next"`
	Cursor	int		`json:"cursor"`
}

type DataResponse struct {
	Data 	interface{}		`json:"data"`
	Paging	Paging			`json:"paging"`	
}