package tools

type ToolInfo interface {
	Name() string
	ExecutableName() string
	GetVersion(tp ToolProvider) (string, error)
	GetVersionWithPath(path string) (string, error)
}
