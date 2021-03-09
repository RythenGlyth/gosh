package shared

// IPluginManager describes the plugin manager.
// The plugin manager is responsible for loading plugins
// and registering their functionality.
type IPluginManager interface {

	// Load loads a plugin from the specified path
	// and returns any encountered error.
	Load(string) error
}
