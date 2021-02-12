package dbquery

import (
	"fmt"
	"gorm.io/gorm"
)

type sortHandler struct{}

func (*sortHandler) Target() queryType {
	return typeSort
}

func (*sortHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	queryObject, ok := value.(*QueryKeyValue)
	if !ok {
		return nil, ErrCastingError
	}

	order, ok := queryObject.value.(Order)
	if !ok {
		return nil, ErrCastingError
	}
	sortStr := fmt.Sprintf("`%s` %s", queryObject.key, order)
	return db.Order(sortStr), nil
}
