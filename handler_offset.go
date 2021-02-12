package dbquery

import (
	"gorm.io/gorm"
)

type offsetHandler struct{}

func (*offsetHandler) Target() queryType {
	return typeOffset
}

func (*offsetHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	result, ok := value.(int)
	if !ok {
		return nil, ErrCastingError
	}
	return db.Offset(result), nil
}
