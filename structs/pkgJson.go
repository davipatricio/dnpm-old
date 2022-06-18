package structs

// From https://docs.npmjs.com/cli/v8/configuring-npm/package-json
type PackageJSONFormat struct {
	Name        string     `json:"name,omitempty"`
	Version     string     `json:"version,omitempty"`
	Description string     `json:"description,omitempty"`
	Keywords    []string   `json:"keywords,omitempty"`
	Homepage    string     `json:"homepage,omitempty"`
	Bugs        BugsFormat `json:"bugs,omitempty"`
	License     string     `json:"license,omitempty"`
	// Can be a string or an array of strings
	Author interface{} `json:"author,omitempty"`
	// Can be a array of strings or an object
	Contributors interface{} `json:"contributors,omitempty"`
	// Can be a string or an object
	Funding interface{} `json:"funding,omitempty"`
	Files   []string    `json:"files,omitempty"`
	Main    string      `json:"main,omitempty"`
	// Can be a string or an array of strings
	Browser interface{}       `json:"browser,omitempty"`
	Bin     map[string]string `json:"bin,omitempty"`
	// Can be a string or an array of strings
	Man         interface{}            `json:"man,omitempty"`
	Directories map[string]interface{} `json:"directories,omitempty"`
	// Can be a string or an object
	Repository           interface{}            `json:"repository,omitempty"`
	Scripts              map[string]string      `json:"scripts,omitempty"`
	Config               map[string]interface{} `json:"config,omitempty"`
	Dependencies         map[string]string      `json:"dependencies,omitempty"`
	PeerDependencies     map[string]string      `json:"peerDependencies,omitempty"`
	PeerDependenciesMeta map[string]interface{} `json:"peerDependenciesMeta,omitempty"`
	DevDependencies      map[string]string      `json:"devDependencies,omitempty"`
	BundleDependencies   []string               `json:"bundleDependencies,omitempty"`
	OptionalDependencies map[string]string      `json:"optionalDependencies,omitempty"`
	Overrides            map[string]interface{} `json:"overrides,omitempty"`
	Engines              map[string]string      `json:"engines,omitempty"`
	Os                   []string               `json:"os,omitempty"`
	Private              bool                   `json:"private,omitempty"`
	Workspaces           []string               `json:"workspaces,omitempty"`
}

type BugsFormat struct {
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}
