package condition_builder

type Builder struct {
	getFieldItem func(string) string
}

func New(getFieldItem func(string) string) Builder {
	return Builder{
		getFieldItem: getFieldItem,
	}
}
