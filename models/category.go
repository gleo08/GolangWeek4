package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (category *Category) TableName() string {
	return "category"
}
