package models

type Review struct {
	Id         int     `json:"id"`
	ProductId  int     `json:"productId"`
	Rating     float32 `json:"rating"`
	CreatedAt  uint32  `json:"createdAt"`
	ModifiedAt uint32  `json:"modifiedAt"`
}

func (r *Review) TableName() string {
	return "review"
}
