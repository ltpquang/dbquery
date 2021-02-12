package dbquery

import "gorm.io/gorm"

type gormQueryExecutor struct {
	db       *gorm.DB
	handlers []queryHandler
}

func (e *gormQueryExecutor) Execute(q Query) (*gorm.DB, error) {
	for _, handler := range e.handlers {
		values, exist := q.data[handler.Target()]
		if !exist {
			continue
		}
		var err error
		for _, value := range values {
			e.db, err = handler.Apply(e.db, value)
			if err != nil {
				return e.db, err
			}
		}
	}
	return e.db, nil
}

func WithGorm(db *gorm.DB) Executor {
	return &gormQueryExecutor{
		db: db,
		handlers: []queryHandler{
			&tableNameHandler{},
			&modelHandler{},
			&preloadHandler{},
			&matchingHandler{},
			&notHandler{},
			&inListHandler{},
			&sortHandler{},
			&countingHandler{},
			&offsetHandler{},
			&limitHandler{},
			&findFirstHandler{},
			&findingHandler{},
		},
	}
}
