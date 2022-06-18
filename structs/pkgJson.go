package structs

// From https://docs.npmjs.com/cli/v8/configuring-npm/package-json
type PackageJSONFormat struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Keywords    []string   `json:"keywords"`
	Homepage    string     `json:"homepage"`
	Bugs        BugsFormat `json:"bugs"`
	License     string     `json:"license"`
	// Can be a string or an array of strings
	Author interface{} `json:"author"`
	// Can be a array of strings or an object
	Contributors interface{} `json:"contributors"`
	// Can be a string or an object
	Funding interface{} `json:"funding"`
	Files   []string    `json:"files"`
	Main    string      `json:"main"`
	// Can be a string or an array of strings
	Browser interface{}       `json:"browser"`
	Bin     map[string]string `json:"bin"`
	// Can be a string or an array of strings
	Man         interface{}            `json:"man"`
	Directories map[string]interface{} `json:"directories"`
	// Can be a string or an object
	Repository           interface{}            `json:"repository"`
	Scripts              map[string]string      `json:"scripts"`
	Config               map[string]interface{} `json:"config"`
	Dependencies         map[string]string      `json:"dependencies"`
	PeerDependencies     map[string]string      `json:"peerDependencies"`
	PeerDependenciesMeta map[string]interface{} `json:"peerDependenciesMeta"`
	DevDependencies      map[string]string      `json:"devDependencies"`
	BundleDependencies   []string               `json:"bundleDependencies"`
	OptionalDependencies map[string]string      `json:"optionalDependencies"`
	Overrides            map[string]interface{} `json:"overrides"`
	Engines              map[string]string      `json:"engines"`
	Os                   []string               `json:"os"`
	Private              bool                   `json:"private"`
	Workspaces           []string               `json:"workspaces"`
}

type BugsFormat struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}
