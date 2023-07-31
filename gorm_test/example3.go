package main

import (
	"fmt"
	"gorm.io/gorm"
)

type condition interface {
	comparable
}

type model interface {
	comparable
}

type UserCond struct {
	ID       uint
	Username string
	Email    string
}

func example3() {
	db := getDB()
	//users := make([]*Users, 0)
	cond := &UserCond{
		ID: 2,
	}

	page := &Page{
		PageIndex: 1,
		Size:      3,
	}

	r, count, _ := findWithCount[*UserCond, *Users](db, cond, page)
	fmt.Println(r)
	fmt.Println(count)
}

func findWithCount[C condition, M model](db *gorm.DB, cond C, page *Page) ([]M, int64, error) {
	data := make([]M, 0)

	model := new(M)
	var count int64
	if err := db.Model(&model).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 查詢物件本體
	if err := db.Offset(page.GetOffset()).
		Limit(page.GetLimit()).Where(&cond).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}
