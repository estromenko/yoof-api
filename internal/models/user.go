package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `form:"email" json:"email" binding:"required" db:"email"`
	Username string `form:"username" json:"username" binding:"required" db:"username"`
	Password string `form:"password" json:"password" binding:"required" password:"password"`
}
