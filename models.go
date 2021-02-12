package dbquery

type queryType string

const (
	typeMatch     = queryType("match")
	typeNot       = queryType("not")
	typeInList    = queryType("in_list")
	typeSort      = queryType("sort")
	typeOffset    = queryType("offset")
	typeLimit     = queryType("limit")
	typeCount     = queryType("count")
	typeFind      = queryType("find")
	typeModel     = queryType("model")
	typePreload   = queryType("preload")
	typeFirst     = queryType("first")
	typeTableName = queryType("table_name")
)

type Order string

const (
	OrderAsc  = Order("ASC")
	OrderDesc = Order("DESC")
)

type Query struct {
	data map[queryType][]interface{}
}

func (q Query) Builder() *Builder {
	return &Builder{
		data: q.data,
	}
}

type operator string

const (
	operatorEqual          = operator("=")
	operatorGreaterOrEqual = operator(">=")
	operatorGreater        = operator(">")
	operatorLessOrEqual    = operator("<=")
	operatorLess           = operator("<")
	operatorNotEqual       = operator("<>")
)

type QueryKeyValue struct {
	key   string
	value interface{}
}

type ComparingObject struct {
	QueryKeyValue
	ope operator
}
