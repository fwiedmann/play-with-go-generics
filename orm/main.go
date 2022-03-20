package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userToCreate := &User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    "Han",
		Surname: "Solo",
	}
	err = CreateGenerics(db, userToCreate)
	if err != nil && err.Error() != "UNIQUE constraint failed: users.id" {
		panic(err.Error())
	}

	searchedUser := &User{}
	err = FindByIdGenerics(db, searchedUser, userToCreate.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(searchedUser)

	searchedUser2 := &User{}
	err = FindById(db, searchedUser2, userToCreate.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(searchedUser2)
}
