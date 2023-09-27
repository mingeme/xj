package api

type PageResponse[T JobData] struct {
	RecordsFiltered int `json:"recordsFiltered"`
	RecordsTotal    int `json:"recordsTotal"`
	Data            []T `json:"data"`
}
