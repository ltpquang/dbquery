package dbquery

import (
	"gorm.io/gorm"
)

type limitHandler struct{}

func (*limitHandler) Target() queryType {
	return typeLimit
}

func (*limitHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	result, ok := value.(int)
	if !ok {
		return nil, ErrCastingError
	}
	return db.Limit(result), nil
}
