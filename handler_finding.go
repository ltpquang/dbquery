package dbquery

import "gorm.io/gorm"

type findingHandler struct{}

func (*findingHandler) Target() queryType {
	return typeFind
}

func (*findingHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	if value == nil {
		return db, nil
	}
	return db.Find(value), nil
}
