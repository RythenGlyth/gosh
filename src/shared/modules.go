package shared

type ModuleIdentifier int

const (
	// ModMain is the id of the main module.
	ModMain ModuleIdentifier = 1

	// ModPrompt is the id of the prompt drawing module.
	ModPrompt ModuleIdentifier = 2

	// ModPluginLoader is the id of the plugin loader.
	ModPluginLoader ModuleIdentifier = 3
)

func (m ModuleIdentifier) AsInt() int {
	return int(m)
}
