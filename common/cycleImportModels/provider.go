package cycleImportModels

import "github.com/jinzhu/gorm"

type Provider struct {
	gorm.Model
	Uin      string
	Nickname string
	Username string
	Gender   int
	Avatar   string
	State    int
	Auths    []Auth `gorm:"foreignkey:ProviderId"`
}

func (Provider) TableName() string {
	return "provider"
}

type Auth struct {
	gorm.Model
	ProviderId   int
	IdentityType int
	Identifier   string
	Credential   string
	State        int
}

func (Auth) TableName() string {
	return "auth"
}
