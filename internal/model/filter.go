package model

type Filter struct {
	Fields map[string]interface{}
	Size   int
}

func NewFilter() *Filter {
	return &Filter{
		Fields: make(map[string]interface{}),
		Size:   0,
	}
}
