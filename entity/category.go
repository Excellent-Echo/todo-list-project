package entity

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	// Todos       []Todo `gorm:"foreignKey:CategoryID"`
}

type CategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
