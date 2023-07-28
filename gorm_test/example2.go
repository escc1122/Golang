package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Data interface {
	comparable
}

type ReturnData[D Data] struct {
	count int64
	Data  []D
}

type Page struct {
	PageIndex int    `gorm:"-"` // 頁碼
	Size      int    `gorm:"-"` // 筆數
	TotalPage int    `gorm:"-"` // 總頁數
	Total     int    `gorm:"-"` // 總筆數
	Order     string `gorm:"-"` // 排序
}

func (p *Page) GetOffset() int {
	return (p.PageIndex - 1) * p.Size
}

func (p *Page) GetLimit() int {
	return p.Size
}

func (p *Page) GetPager() *Page {
	return p
}

func example2() {
	db := getDB()
	users := make([]*Users, 0)
	r, _ := WithPage(db, nil, users)
	fmt.Println(r)
}

func example21() {
	db := getDB()
	users := make([]*Users, 0)

	page := &Page{
		PageIndex: 1,
		Size:      3,
	}

	r, _ := WithPage(db, page, users)
	fmt.Println(r)
}

func WithPage[T Data](db *gorm.DB, pager *Page, r []T) (*ReturnData[T], error) {

	//users := make([]*Users, 0)

	var count int64

	if err := db.Find(&r).Count(&count).Error; err != nil {
		return nil, err
	}

	//pager
	if pager != nil && pager.PageIndex > 0 && pager.Size > 0 {
		if err := db.Find(&r).Offset(pager.GetOffset()).Limit(pager.GetLimit()).Error; err != nil {
			return nil, err
		}

		db = db.Offset(pager.GetOffset()).Limit(pager.GetLimit())

		pager.Total = int(count)
		pager.TotalPage = (int(count) + pager.Size - 1) / pager.Size
	}
	if pager != nil && len(pager.Order) > 0 {
		db = db.Order(pager.Order)
	}

	result := db.Find(&r)

	// 檢查是否有錯誤發生
	if result.Error != nil {
		return nil, result.Error
	}

	returnData := &ReturnData[T]{}
	//member := make([]*po.Member, 0)

	//var count int64

	//if err := db.Find(r, c).Count(&count); err != nil {
	//	return returnData, errors.New("err")
	//}

	returnData.Data = r
	returnData.count = count

	return returnData, nil
}
