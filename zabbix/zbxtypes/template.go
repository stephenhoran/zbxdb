package zbxtypes

type Template struct {
	Sections []Section
}

type Section struct {
	Table     Table
	Fields    []Field
	Unique    Unique
	ChangeLog int
}

type Table struct {
	TableName string
	TableID   string
	TableType string
}

type Field []string

type Unique struct {
	Unique bool
}
