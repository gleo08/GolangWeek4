package models

type Product struct {
	Id            int      `json:"id"`
	CategoryId    int      `json:"categoryId"`
	Name          string   `json:"name"`
	Price         float64  `json:"price"`
	IsSales       bool     `json:"isSales"`
	AverageRating float32  `json:"averageRating" gorm:"-"`
	Reviews       []Review `json:"reviews" gorm:"foreignKey:ProductId`
	Images        []Image  `json:"images" gorm:"foreignKey:ImageId`
	CreatedAt     uint32   `json:"createdAt"`
	ModifiedAt    uint32   `json:"modifiedAt"`
}

func (p *Product) TableName() string {
	return "product"
}
