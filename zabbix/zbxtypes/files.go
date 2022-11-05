package zbxtypes

type ZbxFile string

const (
	SchemaFile     ZbxFile = "templates/schema.tmpl"
	DataFile       ZbxFile = "templates/data.tmpl"
	TemplateFile   ZbxFile = "templates/templates.tmpl"
	DashboardsFile ZbxFile = "templates/dashboard.tmpl"
)
