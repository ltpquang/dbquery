package dbquery

import (
	"fmt"

	"gorm.io/gorm"
)

type matchingHandler struct{}

func (*matchingHandler) Target() queryType {
	return typeMatch
}

func (*matchingHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	matchingObject, ok := value.(*ComparingObject)
	if !ok {
		return nil, ErrCastingError
	}
	queryStr := fmt.Sprintf("%s %s ?", matchingObject.key, matchingObject.ope)
	return db.Where(queryStr, matchingObject.value), nil
}
