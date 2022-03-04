package shared

type AliasManager interface {
	RegisterSimpleAlias(match, expand string)
	RegisterGlobalAlias(match, expand string)
	Expand(line string) string

	ListSimpleAliases() []Alias
	ListGlobalAliases() []Alias
}

type Alias interface {
	Format() string
}
