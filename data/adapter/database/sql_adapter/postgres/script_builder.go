package postgres

import (
	"fmt"
	"strings"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

type ScriptBuilder struct{}

func (ScriptBuilder) buildWhereString(model adapter.ReflectModel, where *query.WhereClause) (string, []any) {
	index := 0
	script := "WHERE "
	values := []any{}

	for {
		index += 1
		script += fmt.Sprintf(`t1."%s" %s $%d`, where.Field, where.Operator, index)
		values = append(values, where.Value)

		if where.Join == nil {
			return script, values
		}

		script += fmt.Sprintf(` %s `, where.JoinType)
		where = where.Join
	}
}

// Build a select query using the reflection model.
// Note that the model's name and fields are obtained using reflection and therefore sql injection is not possible.
func (s ScriptBuilder) BuildSelectQuery(model adapter.ReflectModel, where *query.WhereClause) (string, []any) {
	script := `
		SELECT t1."%s"
			FROM "%s" t1
	`
	values := []any{}

	if where != nil {
		s, v := s.buildWhereString(model, where)
		script += s
		values = v
	}

	return fmt.Sprintf(
		script, strings.Join(model.Fields, `", t1."`), model.Name,
	), values
}

// Build an insert statement using the reflection model.
// Note that the model's name and fields are obtained using reflection and therefore sql injection is not possible.
func (s ScriptBuilder) BuildInsertStatement(model adapter.ReflectModel) string {
	script := `
		INSERT INTO "%s" ("%s")
			VALUES (%s)
	`

	return fmt.Sprintf(
		script, model.Name, strings.Join(model.Fields, `", "`), s.buildParametrizedString(model),
	)
}

func (ScriptBuilder) buildParametrizedString(model adapter.ReflectModel) string {
	params := []string{}
	for index := range model.Fields {
		params = append(params, fmt.Sprintf("$%d", index+1))
	}
	return strings.Join(params, ", ")
}

// Build an update statement using the reflection model.
// Note that the model's name and fields are obtained using reflection and therefore sql injection is not possible.
func (s ScriptBuilder) BuildUpdateStatement(model adapter.ReflectModel) string {
	script := `
		UPDATE "%s" SET
			%s
		WHERE "%s" = $1
	`

	return fmt.Sprintf(
		script, model.Name, s.buildSetString(model), model.UniqueField(),
	)
}

func (ScriptBuilder) buildSetString(model adapter.ReflectModel) string {
	setScripts := []string{}
	for index, field := range model.Fields {
		if index == 0 {
			continue
		}

		setScripts = append(setScripts,
			fmt.Sprintf(`"%s" = $%d`, field, index+1),
		)
	}

	return strings.Join(setScripts, ", ")
}

// Build a delete statement using the reflection model.
// Note that the model's name and fields are obtained using reflection and therefore sql injection is not possible.
func (s ScriptBuilder) BuildDeleteStatement(model adapter.ReflectModel, where *query.WhereClause) (string, []any) {
	script := `
		DELETE FROM "%s" t1
	`
	values := []any{}

	if where != nil {
		s, v := s.buildWhereString(model, where)
		script += s
		values = v
	}

	return fmt.Sprintf(
		script, model.Name,
	), values
}
