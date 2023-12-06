package model

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

type Rooms struct {
	ID           string `gorm:"unique"`
	Name         string
	PhotoProfile string
	LastMessage  string
}

type Message struct {
	ID           uint `gorm:"primaryKey"`
	RoomID       string
	Message      string
}
