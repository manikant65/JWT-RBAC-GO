package model

type Login struct {
	Username string `json:"username" binding:"required"` // to safely pass user input into sql query
	Password string `json:"password" binding:"required"`
}
