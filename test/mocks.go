package test

import "go_web_server/internal/model"

var BookMock = model.Book{
	ID:     1,
	Name:   "Mafia",
	Genre:  "fantasy",
	Author: "Roman",
}

const userAge = 18

var UserMock = model.User{
	ID:   1,
	Name: "Roma",
	Age:  userAge,
	City: "Toronto",
}
