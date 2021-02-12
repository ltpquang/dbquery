package dbquery

import "gorm.io/gorm"

type Executor interface {
	Execute(q Query) (*gorm.DB, error)
}

type queryHandler interface {
	Target() queryType
	Apply(db *gorm.DB, value interface{}) (*gorm.DB, error)
}
