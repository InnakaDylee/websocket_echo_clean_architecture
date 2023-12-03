package websocket

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email	string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}