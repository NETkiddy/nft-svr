package cycleImportModels

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	OpenID       string        `json:"openId"`
	Name         string        `json:"name"`
	Nickname     string        `json:"nickName"`
	Gender       int           `json:"gender"`
	Language     string        `json:"language"`
	City         string        `json:"city"`
	Province     string        `json:"province"`
	Country      string        `json:"country"`
	AvatarURL    string        `json:"avatarUrl"`
	Timestamp    int           `json:"timestamp"`
	Appid        string        `json:"appid"`
	Products     []Product     `gorm:"many2many:user_product;association_jointable_foreignkey:product_id;jointable_foreignkey:user_id;"`
	UserProducts []UserProduct `gorm:"foreignkey:UserId"`
}

func (User) TableName() string {
	return "user"
}

type UserProduct struct {
	ID          uint `gorm:"primary_key"`
	User        User
	UserId      uint
	Product     Product
	ProductId   uint
	Price       int
	Description string
}

type UserProductAlso struct {
	gorm.Model
	UserId      uint
	ProductId   uint
	Description string
}
func (UserProductAlso) TableName() string {
	return "user_product"
}