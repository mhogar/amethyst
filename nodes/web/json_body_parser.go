package web

import (
	"encoding/json"
	"errors"

	"github.com/mhogar/kiwi/nodes"
)

type JSONBodyParserNode[Model any] struct{}

func NewJSONBodyParserNode[Model any]() JSONBodyParserNode[Model] {
	return JSONBodyParserNode[Model]{}
}

func (JSONBodyParserNode[Model]) Run(ctx interface{}, _ any) (any, *nodes.Error) {
	model := new(Model)

	decoder := json.NewDecoder(ctx.(WebContext).GetRequest().Body)
	err := decoder.Decode(model)
	if err != nil {
		return nil, nodes.ClientError(errors.New("invalid json body"))
	}

	return model, nil
}
