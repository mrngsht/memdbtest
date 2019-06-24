package memdbtest

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID       int    `gorm:"AUTO_INCREMENT;primary_key"`
	Category string `gorm:"not null;size:8"`
	Content  string `gorm:"not null;size:255"`
}

func Create(db *gorm.DB, category string, content string) error {
	return db.Create(&Todo{Category: category, Content: content}).Error
}

func Read(db *gorm.DB, todo *Todo, id int) error {
	return db.Find(todo, "id = ?", id).Error
}

func Update(db *gorm.DB, id int, category string, content string) error {
	return db.Model(&Todo{}).Where("id = ?", id).Updates(&Todo{Category: category, Content: content}).Error
}

func Delete(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).Delete(&Todo{}).Error
}
