package memdbtest

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestCreate(t *testing.T) {
	db := createInMemoryDb(t)
	defer db.Close()
	fatalIfError(t, db.AutoMigrate(&Todo{}).Error)

	fatalIfError(t, Create(db, "first", "first todo"))

	var todo Todo
	fatalIfError(t, db.First(&todo).Error)
	assertEqual(t, "first", todo.Category)
	assertEqual(t, "first todo", todo.Content)
}

func TestRead(t *testing.T) {
	db := createInMemoryDb(t)
	defer db.Close()
	fatalIfError(t, db.AutoMigrate(&Todo{}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "first", Content: "first todo"}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "second", Content: "second todo"}).Error)

	var todo Todo
	fatalIfError(t, Read(db, &todo, 1))

	assertEqual(t, "first", todo.Category)
	assertEqual(t, "first todo", todo.Content)
}

func TestUpdate(t *testing.T) {
	db := createInMemoryDb(t)
	defer db.Close()
	fatalIfError(t, db.AutoMigrate(&Todo{}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "first", Content: "first todo"}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "second", Content: "second todo"}).Error)

	fatalIfError(t, Update(db, 1, "first:updated", "first todo:updated"))

	var first, second Todo
	fatalIfError(t, db.Find(&first, "id = 1").Error)
	fatalIfError(t, db.Find(&second, "id = 2").Error)
	assertEqual(t, "first:updated", first.Category)
	assertEqual(t, "first todo:updated", first.Content)
	assertEqual(t, "second", second.Category)
}

func TestDelete(t *testing.T) {
	db := createInMemoryDb(t)
	defer db.Close()
	fatalIfError(t, db.AutoMigrate(&Todo{}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "first", Content: "first todo"}).Error)
	fatalIfError(t, db.Create(&Todo{Category: "second", Content: "second todo"}).Error)

	fatalIfError(t, Delete(db, 1))

	var count int
	var todo Todo
	fatalIfError(t, db.Table("todos").Count(&count).Error)
	fatalIfError(t, db.First(&todo).Error)

	assertEqual(t, 1, count)
	assertEqual(t, "second", todo.Category)
}

func createInMemoryDb(t *testing.T) *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("db open error: %v", err)
	}
	return db
}

func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("fatal: %v\n", err)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("test failure. expected %v, actual: %v\n", expected, actual)
	}
}