package models

type Image struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
}

func (i *Image) TableName() string {
	return "image"
}
