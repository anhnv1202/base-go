package po

var models []any

func Register(model any) {
	models = append(models, model)
}

func All() []any {
	return models
}
