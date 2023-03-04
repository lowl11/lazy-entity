package entity_models

type WhereCondition struct {
	Left  string
	Right []any

	Equals  bool
	More    bool
	Less    bool
	Like    bool
	ILike   bool
	Between bool
}
