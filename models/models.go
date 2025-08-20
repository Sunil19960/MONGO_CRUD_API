package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
