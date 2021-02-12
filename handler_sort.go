package dbquery

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: queryObject.key},
			Desc:   order == OrderDesc}),
		nil
}
