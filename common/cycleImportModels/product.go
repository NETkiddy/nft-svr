package cycleImportModels

import "github.com/jinzhu/gorm"

type TokenClasses struct {
	gorm.Model
	Name          string `json:"name"`
	Description   string `json:"description"`
	Total         string `json:"total"`
	Renderer      string `json:"renderer"`
	CoverImageUrl string `json:"cover_image_url"`
}

func (TokenClasses) TableName() string {
	return "token_classes"
}
