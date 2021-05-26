package entity

type Category struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
