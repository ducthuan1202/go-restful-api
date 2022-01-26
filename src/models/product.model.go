package models

// Product is struct
type Product struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	CategoryID uint   `gorm:"not null" json:"category_id"`

	Category Category `json:"category,omitempty"`
}
