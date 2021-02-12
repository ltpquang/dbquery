package dbquery

import (
	"gorm.io/gorm"
)

type modelHandler struct{}

func (*modelHandler) Target() queryType {
	return typeModel
}

func (*modelHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	if value == nil {
		return db, nil
	}
	return db.Model(value), nil
}
