package dbquery

type Builder struct {
	data map[queryType][]interface{}
}

func NewBuilder() *Builder {
	b := &Builder{data: make(map[queryType][]interface{})}
	return b
}

func (b *Builder) Build() Query {
	return Query{data: b.data}
}

func (b *Builder) add(qType queryType, object interface{}) *Builder {
	b.ensure(qType)
	b.data[qType] = append(b.data[qType], object)
	return b
}

func (b *Builder) ensure(qType queryType) {
	if _, exist := b.data[qType]; !exist {
		b.data[qType] = make([]interface{}, 0)
	}
}

func (b *Builder) TableName(tableName string) *Builder {
	return b.add(typeTableName, tableName)
}

func (b *Builder) Model(model interface{}) *Builder {
	return b.add(typeModel, model)
}

func (b *Builder) Preload(column string, condition ...interface{}) *Builder {
	var cdt []interface{}
	if len(condition) != 0 {
		cdt = condition
	}
	return b.add(typePreload, &QueryKeyValue{
		key:   column,
		value: cdt,
	})
}

func (b *Builder) Find(findingResult interface{}) *Builder {
	return b.add(typeFind, findingResult)
}

func (b *Builder) First(firstResult interface{}) *Builder {
	return b.add(typeFirst, firstResult)
}

func (b *Builder) Count(countingResult *int64) *Builder {
	return b.add(typeCount, countingResult)
}

func (b *Builder) Sort(fieldName string, order Order) *Builder {
	return b.add(typeSort, &QueryKeyValue{
		key:   fieldName,
		value: order,
	})
}

func (b *Builder) Offset(offset int) *Builder {
	return b.add(typeOffset, offset)
}

func (b *Builder) Limit(limit int) *Builder {
	return b.add(typeLimit, limit)
}

func (b *Builder) InList(fieldName string, list []interface{}) *Builder {
	return b.add(typeInList, &QueryKeyValue{
		key:   fieldName,
		value: list,
	})
}

func (b *Builder) Not(fieldName string, value interface{}) *Builder {
	return b.add(typeNot, &QueryKeyValue{
		key:   fieldName,
		value: value,
	})
}

func (b *Builder) Equal(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorEqual, value)
}

func (b *Builder) GreaterOrEqual(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorGreaterOrEqual, value)
}

func (b *Builder) Greater(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorGreater, value)
}

func (b *Builder) LessOrEqual(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorLessOrEqual, value)
}

func (b *Builder) Less(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorLess, value)
}

func (b *Builder) NotEqual(fieldName string, value interface{}) *Builder {
	return b.match(fieldName, operatorNotEqual, value)
}

func (b *Builder) match(fieldName string, ope operator, value interface{}) *Builder {
	return b.add(typeMatch, &ComparingObject{
		QueryKeyValue: QueryKeyValue{
			key:   fieldName,
			value: value,
		},
		ope: ope,
	})
}
