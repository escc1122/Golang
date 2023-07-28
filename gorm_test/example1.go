package main

import "fmt"

func example1() {
	db := getDB()
	users := make([]*Users, 0)
	result := db.Find(&users)

	if result.Error != nil {
		panic("failed to query users")
	}
	fmt.Println(users)
}
