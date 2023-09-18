package types

type Version struct {
	// Version indicates the version in SemVer format
	Version string `yaml:"version"`

	// Vars indicates a key-value pair of variables that will be used
	// for rendering the HTTP request.
	Vars map[string]string `yaml:"vars"`
}
