package dbquery

import (
	"gorm.io/gorm"
)

type tableNameHandler struct{}

func (*tableNameHandler) Target() queryType {
	return typeTableName
}

func (*tableNameHandler) Apply(db *gorm.DB, value interface{}) (*gorm.DB, error) {
	if value == nil {
		return db, nil
	}
	tableName, ok := value.(string)
	if !ok {
		return nil, ErrCastingError
	}
	if len(tableName) == 0 {
		return db, nil
	}
	return db.Table(tableName), nil
}
