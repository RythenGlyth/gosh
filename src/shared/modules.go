package shared

import "strconv"

type ModuleIdentifier int

const (
	// ModMain is the id of the main module.
	ModMain ModuleIdentifier = 1

	// ModPrompt is the id of the prompt drawing module.
	ModPrompt ModuleIdentifier = 2

	// ModPluginLoader is the id of the plugin loader.
	ModPluginLoader ModuleIdentifier = 3
)

func ModuleIdentifierFromInt(m int) ModuleIdentifier {
	return ModuleIdentifier(m)
}

func (m ModuleIdentifier) AsInt() int {
	return int(m)
}

func (m ModuleIdentifier) String() string {
	switch m {
	case ModMain:
		return "  main"
	case ModPrompt:
		return "prompt"
	case ModPluginLoader:
		return "loader"
	default:
		return strconv.Itoa(int(m))
	}
}
