package dbquery

type Operator string

const (
	Equal          = Operator(" = ?")
	Greater        = Operator(" >= ?")
	GreaterOrEqual = Operator(" > ?")
	LessOrEqual    = Operator(" <= ?")
	Less           = Operator(" < ?")
	NotEqual       = Operator(" != ?")
)

type Values struct {
	Operator Operator
	Value    interface{}
}

func New(value interface{}) *Values {
	return &Values{
		Operator: Equal,
		Value:    value,
	}
}
func NewWithOperator(operator Operator, value interface{}) *Values {
	return &Values{
		Operator: operator,
		Value:    value,
	}
}
