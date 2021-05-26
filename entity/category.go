package entity

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	// Todos       []Todo `gorm:"foreignKey:CategoryID"`
}
