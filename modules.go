package main

import "time"

type Person struct {
	Firstname string `json:"firstname"   binding:"required"`
	Lastname  string `json:"lastname"   binding:"required"`
}

type Content struct {
	Title string `json:"title"   binding:"required"`
	Body  string `json:"body"   binding:"required"`
}

type Article struct {
	Id     string `json:"id" `
	Author Person `json:"Muallif" binding:"required" `
	Content
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt *time.Time `json:"updated_at" `
}
