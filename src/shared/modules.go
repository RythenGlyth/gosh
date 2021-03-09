package shared

type ModuleIdentifier int

const (
	// ModMain is the id of the main module.
	ModMain ModuleIdentifier = iota

	// ModPrompt is the id of the prompt drawing module.
	ModPrompt ModuleIdentifier = iota

	// ModPluginLoader is the id of the plugin loader.
	ModPluginLoader ModuleIdentifier = iota
)

func (m ModuleIdentifier) AsInt() int {
	return int(m)
}
