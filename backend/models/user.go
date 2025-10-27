package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	IpAddress string `json:"ip_address"`
}
