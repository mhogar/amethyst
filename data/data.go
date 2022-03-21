package data

type WhereClause struct {
	Field    string
	Operator string
	value    interface{}
}

func Where(field string, op string, val interface{}) WhereClause {
	return WhereClause{
		Field:    field,
		Operator: op,
		value:    val,
	}
}

func Select[T any](_ ...WhereClause) *T {
	var model T
	// fill in fields based on query
	return &model
}
