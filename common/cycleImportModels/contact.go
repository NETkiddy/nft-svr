package cycleImportModels

import "github.com/jinzhu/gorm"

type Contact struct {
	gorm.Model
	AccountId            uint
	ContactUuid          string
	Name                 string
	Company              string
	Email                string
	WxId                 string
	Phone                string
	CompanyPhone         string
	Title                string
	Website              string
	Address              string
	Contactgroups        []Contactgroup        `gorm:"many2many:contactgroup_contact;association_jointable_foreignkey:contactkgroup_id;jointable_foreignkey:contact_id;"`
	ContactgroupContacts []ContactgroupContact `gorm:"foreignkey:ContactId"`
}

func (Contact) TableName() string {
	return "contact"
}

type Contactgroup struct {
	gorm.Model
	AccountId            uint
	ContactgroupUuid     string
	Name                 string
	Description          string
	Contacts             []Contact             `gorm:"many2many:contactgroup_contact;association_jointable_foreignkey:contact_id;jointable_foreignkey:contactgroup_id;"`
	ContactgroupContacts []ContactgroupContact `gorm:"foreignkey:ContactgroupId"`
}

func (Contactgroup) TableName() string {
	return "contactgroup"
}

type ContactgroupContact struct {
	ID             uint `gorm:"primary_key"`
	Contactgroup   Contactgroup
	ContactgroupId uint
	Contact        Contact
	ContactId      uint
}

func (ContactgroupContact) TableName() string {
	return "contactgroup_contact"
}
