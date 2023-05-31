package impl

import (
	"quocbang/golang-to-do-list/impl/orm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataManager struct {
	db *gorm.DB
}

func NewDB(schema string) (*gorm.DB, error) {
	dsn := "host=198.1.1.92 user=kendakv password=kenda dbname=kverp port=5432 sslmode=disable search_path=test"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	maybeMigrateTables(db)
	return db, nil
}

// maybeMigrateTables attempts to create tables automatically if implement
// models.Model interface.
func maybeMigrateTables(db *gorm.DB) error {
	ms := models.GetModelList()
	dst := []any{}
	for _, m := range ms {
		dst = append(dst, m)
	}
	return db.AutoMigrate(dst...)
}
