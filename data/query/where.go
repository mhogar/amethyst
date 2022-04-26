package query

type WhereClause struct {
	Field    string
	Operator string
	Value    interface{}

	Join     *WhereClause
	JoinType string
}

func (w *WhereClause) And(join *WhereClause) *WhereClause {
	w.Join = join
	w.JoinType = "AND"
	return w
}

func (w *WhereClause) Or(join *WhereClause) *WhereClause {
	w.Join = join
	w.JoinType = "OR"
	return w
}

func Where(field string, op string, val interface{}) *WhereClause {
	return &WhereClause{
		Field:    field,
		Operator: op,
		Value:    val,
	}
}
