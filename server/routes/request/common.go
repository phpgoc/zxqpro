package request

import "errors"

type Page struct {
	Page     int `json:"page" form:"page" binding:"required,min=1" default:"1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,min=5" default:"10"`
}

type CommonID struct {
	ID uint `json:"id" form:"id" binding:"required,min=1"`
}

type Order byte

// OrderBy 排序规则结构体
type OrderBy struct {
	// 排序字段，必填
	Field string `json:"field" form:"field" binding:"required"`
	// 排序方向，必填
	Desc bool `json:"desc" form:"desc" binding:"required"`
}

func IsAllValidOrder(order []OrderBy, validList []string) error {
out:
	for _, o := range order {
		for _, v := range validList {
			if o.Field == v {
				continue out
			}
		}
		return errors.New("排序字段不合法" + o.Field)
	}
	return nil
}
