package dbquery

import "gorm.io/gorm"

type findFirstHandler struct{}

func (*findFirstHandler) Target() queryType {
	return typeFirst
}

func (*findFirstHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	if value == nil {
		return db, nil
	}
	return db.First(value), nil
}
