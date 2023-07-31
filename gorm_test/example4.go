package main

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

func findWithCount2(db *gorm.DB, page *Page, cond any, data any) (int64, error) {
	var count int64
	if data == nil {
		return 0, errors.New("data is nil")
	}

	dataType := reflect.TypeOf(data)

	for dataType.Kind() != reflect.Pointer {
		return 0, errors.New("data is not Pointer")
	}

	if cond != nil {
		db = db.Where(cond)
	}

	if err := db.Model(data).Count(&count).Error; err != nil {
		return 0, err
	}

	if page != nil {
		db = db.Offset(page.GetOffset()).Limit(page.GetLimit())
	}

	// 查詢物件本體
	if err := db.Find(data).Error; err != nil {
		return 0, err
	}

	return count, nil
}
