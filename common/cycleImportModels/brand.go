package cycleImportModels

import "github.com/jinzhu/gorm"

type Brand struct {
	gorm.Model
	ProviderId     uint
	Name           string
	Picurl         string
	DetailImg      string
	CategoryLarge  string
	CategoryMedium string
	CategorySmall  string
	Tags           string
	Homepage       string
	Description    string
	State          int
	Products       []Product      `gorm:"many2many:brand_product;association_jointable_foreignkey:product_id;jointable_foreignkey:brand_id;"`
	BrandProducts  []BrandProduct `gorm:"foreignkey:BrandId"`
}

func (Brand) TableName() string {
	return "brand"
}

type BrandProduct struct {
	ID        uint `gorm:"primary_key"`
	Brand     Brand
	BrandId   uint
	Product   Product
	ProductId uint
}

func (BrandProduct) TableName() string {
	return "brand_product"
}
