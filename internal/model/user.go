package model

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  uint   `json:"age" binding:"required"`
	City string `json:"city" binding:"required"`
}
