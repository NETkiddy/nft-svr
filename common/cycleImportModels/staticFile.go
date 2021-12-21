package cycleImportModels

import "github.com/jinzhu/gorm"

type StaticFile struct {
	gorm.Model
}

func (StaticFile) TableName() string {
	return "static_file"
}
