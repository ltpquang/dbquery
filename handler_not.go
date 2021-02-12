package dbquery

import (
	"gorm.io/gorm"
)

type notHandler struct{}

func (*notHandler) Target() queryType {
	return typeNot
}

func (*notHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	kvObject, ok := value.(*QueryKeyValue)
	if !ok {
		return nil, ErrCastingError
	}
	return db.Not(kvObject.key, kvObject.value), nil
}
