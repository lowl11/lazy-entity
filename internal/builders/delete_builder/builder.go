package delete_builder

type Builder struct {
	query string
}

func New() *Builder {
	return &Builder{}
}
