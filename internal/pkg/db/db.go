package db

import (
	"fmt"
	"os"

	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func Prepare(logger logger.Interface) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func All[T any](db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var records []T
	err := db.Scopes(scopes...).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func First[T any](db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (T, error) {
	var record T
	err := db.Scopes(scopes...).First(&record).Error
	if err != nil {
		return record, err
	}
	return record, nil
}

func CreateOrFirst[T any](db *gorm.DB, pointer *T, findScope func(*gorm.DB) *gorm.DB) error {
	err := db.Create(pointer).Error

	if !IsErrorDuplicateEntry(err) {
		return err
	}

	return db.Scopes(findScope).First(pointer).Error

}
func CreateOrUpdate[T any](db *gorm.DB, pointer *T, findScope func(*gorm.DB) *gorm.DB) error {
	updateValue := *pointer // NOTE: Get snapshot before create
	err := db.Create(pointer).Error

	if !IsErrorDuplicateEntry(err) {
		return err
	}

	if err := db.Scopes(findScope).First(&pointer).Error; err != nil {
		return err
	}

	return db.Model(pointer).Updates(updateValue).Error
}

func Count[T any](db *gorm.DB, findScope func(*gorm.DB) *gorm.DB) (int, error) {
	var cou int64
	var record T
	err := db.Model(&record).Scopes(findScope).Count(&cou).Error
	if err != nil {
		return 0, err
	}
	return int(cou), nil
}

func LockForUpdate[T any](db *gorm.DB, findScope func(*gorm.DB) *gorm.DB) (T, error) {
	return First[T](db.Clauses(clause.Locking{Strength: "UPDATE"}), findScope)
}
func IsErrorDuplicateEntry(err error) bool {
	switch v := err.(type) {
	case *driver.MySQLError:
		switch v.Number {
		case 1062:
			return true
		}
	}
	return false
}
