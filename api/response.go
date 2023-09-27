package api

type BaseResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Content string `json:"content"`
}

type PageResponse[T JobData | GroupData] struct {
	RecordsFiltered int `json:"recordsFiltered"`
	RecordsTotal    int `json:"recordsTotal"`
	Data            []T `json:"data"`
}
