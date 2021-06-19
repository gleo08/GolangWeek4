package models

type Review struct {
	Id         int     `json:"id"`
	ProductId  int     `json:"productId"`
	Rating     float32 `json:"rating"`
	CreatedAt  int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}

func (r *Review) TableName() string {
	return "review"
}
