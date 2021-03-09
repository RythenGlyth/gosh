package shared

// IPluginManager describes the
type IPluginManager interface {

	// Load loads a plugin from the specified path
	// and returns any encountered error.
	Load(string) error
}
