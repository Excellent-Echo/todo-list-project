package migration

import "time"

type Todo struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	CategoryID  int
	UserID      int
	IsComplete  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

type Category struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Todos       []Todo `gorm:"foreignKey:CategoryID"`
}

type UserDetail struct {
	ID          int `gorm:"primaryKey"`
	NoHandphone uint
	Gender      string
	Address     string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
}

type User struct {
	ID         int          `gorm:"primaryKey" json:"id"`
	FirstName  string       `json:"first_name"`
	LastName   string       `json:"last_name"`
	Email      string       `json:"email"`
	Password   string       `json:"-"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  time.Time    `gorm:"index" json:"-"`
	Todos      []Todo       `gorm:"foreignKey:UserID"`
	UserDetail []UserDetail `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"` // kita tangkap dari file (foto) , path / dir file foto
	UserID      int    `json:"user_id"`
}
