package models

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Birthday   string `json:"birthday"`
	Status     bool   `json:"status"`
	CreatedAt  uint32 `json:"createdAt"`
	ModifiedAt uint32 `json:"modifiedAt"`
}

func (u *User) TableName() string {
	return "user"
}
