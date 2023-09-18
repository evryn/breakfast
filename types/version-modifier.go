package types

type VersionModifier struct {
	// Regex indicates whether the modifier should apply to a version or not
	Format string `yaml:"format"`

	// Vars indicates a key-value pair of variables that will be used
	// for rendering the HTTP request.
	Vars map[string]string `yaml:"vars"`
}
