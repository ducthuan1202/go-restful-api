package models

// Category is struct
type Category struct {

	// them mixin field (id, created_at, updated_at, deleted_at)
	// gorm.Model

	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`

	Products []Product `gorm:"foreignkey:CategoryID" json:"products,omitempty"`
}
