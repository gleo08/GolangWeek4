package models

type Cart struct {
	Id          int     `json:"id"`
	CategoryId  int     `json:"categoryId"`
	Image       string  `json:"image"`
	ProductName string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   uint32  `json:"createdAt"`
	ModifiedAt  uint32  `json:"modifiedAt"`
}

func (cart *Cart) TableName() string {
	return "cart"
}
