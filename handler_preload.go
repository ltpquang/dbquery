package dbquery

import (
	"gorm.io/gorm"
)

type preloadHandler struct{}

func (*preloadHandler) Target() queryType {
	return typePreload
}

func (p *preloadHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	queryObject, ok := value.(*QueryKeyValue)
	if !ok {
		return nil, ErrCastingError
	}
	conditions, err := p.getConditions(queryObject)
	if err != nil {
		return nil, err
	}
	if len(conditions) != 0 {
		return db.Preload(queryObject.key, conditions...), nil
	} else {
		return db.Preload(queryObject.key), nil
	}
}

func (*preloadHandler) getConditions(queryObject *QueryKeyValue) ([]interface{}, error) {
	if queryObject.value == nil {
		return nil, nil
	}
	inputCondts, ok := queryObject.value.([]interface{})
	if !ok {
		return nil, ErrCastingError
	}
	return inputCondts, nil
}
