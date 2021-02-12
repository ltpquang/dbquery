package dbquery

import (
	"fmt"
	"gorm.io/gorm"
)

type inListHandler struct{}

func (*inListHandler) Target() queryType {
	return typeInList
}

func (*inListHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	queryObject, ok := value.(*QueryKeyValue)
	if !ok {
		return nil, ErrCastingError
	}
	list, ok := queryObject.value.([]interface{})
	if !ok {
		return nil, ErrCastingError
	}
	queryStr := fmt.Sprintf("%s in (?)", queryObject.key)
	return db.Where(queryStr, list), nil
}
