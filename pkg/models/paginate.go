package models

type Paginate struct {
	Offset  int    `query:"offset"`
	Limit   int    `query:"limit"`
	Total   int64  `query:"total"`
	Keyword string `query:"keyword"`
}
