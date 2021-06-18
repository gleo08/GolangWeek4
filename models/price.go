package models

type Price struct {
	Id         int     `json:"id"`
	ProductId  int     `json:"product_id"`
	Value      float64 `json:"value"`
	CreatedAt  uint32  `json:"createdAt"`
	ModifiedAt uint32  `json:"modifiedAt"`
}

func (p *Price) TableName() string {
	return "price"
}
