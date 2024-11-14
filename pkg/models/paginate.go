package models

import "time"

type Paginate struct {
	Offset  int    `query:"offset"`
	Limit   int    `query:"limit"`
	Total   int64  `query:"total"`
	Keyword string `query:"keyword"`
}
type DatePicker struct {
	From time.Time `query:"from"`
	To   time.Time `query:"to"`
}
