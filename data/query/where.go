package query

const (
	WHERE_JOIN_TYPE_AND = iota
	WHERE_JOIN_TYPE_OR  = iota
)

type WhereClause struct {
	Field    string
	Operator string
	Value    interface{}

	Join     *WhereClause
	JoinType int
}

func (w *WhereClause) And(join *WhereClause) *WhereClause {
	w.Join = join
	w.JoinType = WHERE_JOIN_TYPE_AND
	return w
}

func (w *WhereClause) Or(join *WhereClause) *WhereClause {
	w.Join = join
	w.JoinType = WHERE_JOIN_TYPE_OR
	return w
}

func Where(field string, op string, val interface{}) *WhereClause {
	return &WhereClause{
		Field:    field,
		Operator: op,
		Value:    val,
		Join:     nil,
	}
}
