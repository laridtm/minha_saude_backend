package model

import "time"

type Filter struct {
	Fields   map[string]interface{}
	FromDate *time.Time
	Size     int
}

func NewFilter() *Filter {
	return &Filter{
		Fields:   make(map[string]interface{}),
		FromDate: nil,
		Size:     0,
	}
}
