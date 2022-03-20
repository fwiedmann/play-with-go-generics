package main

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

func isPointer(data any) error {
	if reflect.ValueOf(data).Type().Kind() != reflect.Pointer {
		return errors.New("is not a pointer")
	}
	return nil
}

func FindById(db *gorm.DB, data any, id uint) error {
	if err := isPointer(data); err != nil {
		return err
	}

	if err := db.AutoMigrate(data); err != nil {
		return err
	}
	return db.Find(data, "ID = ?", id).Error
}

func FindByIdGenerics[T any](db *gorm.DB, data T, id uint) error {
	if err := isPointer(data); err != nil {
		return err
	}

	if err := db.AutoMigrate(data); err != nil {
		return err
	}
	return db.Find(data, "ID = ?", id).Error
}

func CreateGenerics[T any](db *gorm.DB, data T) error {

	if err := isPointer(data); err != nil {
		return err
	}

	if err := db.AutoMigrate(data); err != nil {
		return err
	}
	return db.Create(data).Error
}
