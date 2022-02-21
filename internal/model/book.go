package model

type Book struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
	Author string `json:"author" binding:"required"`
}
