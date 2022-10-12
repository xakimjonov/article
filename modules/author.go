package modules

import "time"

type Author struct {
	Id        string     `json:"id" `
	Firstname string     `json:"firstname"   binding:"required"  minLength:"2" maxLength:"30" example:"Jack"`
	Lastname  string     `json:"lastname"   binding:"required" minLength:"2" maxLength:"30" example:"Haldson"`
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" `
}

type MakeAuthor struct {
	Firstname string     `json:"firstname"   binding:"required"  minLength:"2" maxLength:"30" example:"Jack"`
	Lastname  string     `json:"lastname"   binding:"required" minLength:"2" maxLength:"30" example:"Haldson"`
	
}

