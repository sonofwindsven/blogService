package model

import (
	"blogService/pkg/app"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

// 仅仅用于文档返回
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	db.Where("state = ?", t.State)
	if t.Name != "" {
		db.Where("name = ?", t.Name)
	}
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	db.Where("state = ?", t.State)
	if t.Name != "" {
		db.Where("name = ?", t.Name)
	}
	db.Offset(pageOffset).Limit(pageSize)
	if err := db.Model(&t).Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del = ?", t.ID, 0).Update(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.ID, 0).Delete(&t).Error
}
