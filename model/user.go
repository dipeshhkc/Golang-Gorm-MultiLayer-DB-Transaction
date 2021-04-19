package model

import "gorm.io/gorm"

// User ... User Database Model
type User struct {
	gorm.Model
	Email  string  `json:"email" gorm:"unique;not null"`
	Wallet float64 `json:"wallet"`
}

//MoneyTransfer --- MoneyTransfer Struct
type MoneyTransfer struct {
	Receiver uint    `json:"receiver"`
	Giver    uint    `json:"giver"`
	Amount   float64 `json:"amount"`
}
