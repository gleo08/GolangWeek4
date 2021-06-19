package models

type Price struct {
	Id         int     `json:"id"`
	ProductId  int     `json:"product_id"`
	Value      float64 `json:"value"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}

func (p *Price) TableName() string {
	return "price"
}
