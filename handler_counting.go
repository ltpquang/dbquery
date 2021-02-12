package dbquery

import (
	"gorm.io/gorm"
)

type countingHandler struct{}

func (*countingHandler) Target() queryType {
	return typeCount
}

func (*countingHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	if value == nil {
		return db, nil
	}
	result, ok := value.(*int64)
	if !ok {
		return nil, ErrCastingError
	}
	return db.Count(result), nil
}
