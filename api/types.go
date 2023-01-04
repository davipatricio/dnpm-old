package api

type PackageJSONTypeField string

const (
	PackageJSONTypeModule   PackageJSONTypeField = "module"
	PackageJSONTypeCommonJS PackageJSONTypeField = "commonjs"
)

type PackageJSONBugsField struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

// Can be a string or an object
type PackageJSONPersonField any

type PackageJSONPersonFieldObject struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

// Can be a string, URL, or an array of those
type PackageJSONFundingField any

// Can be a string, URL, or an array of those
type PackageJSONFundingFieldArray []any

type PackageJSONFundingFieldObject struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Can be a string or an object
type PackageJSONRepositoryField any

type PackageJSONRepositoryFieldObject struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type PackageJSON struct {
	// The name of the package
	Name string `json:"name"`
	// The version of the package
	Version string `json:"version"`
	// Put a description in it. It's a string. This helps people discover your package
	Description string `json:"description"`
	// Put keywords in it. It's an array of strings. This helps people discover your package
	Keywords []string `json:"keywords"`
	// The url to the project homepage
	Homepage string `json:"homepage"`
	// The url to your project's issue tracker and / or the email address to which issues should be reported. These are helpful for people who encounter issues with your package
	Bugs          PackageJSONBugsField     `json:"bugs"`
	Author        PackageJSONPersonField   `json:"author"`
	Contribuitors []PackageJSONPersonField `json:"contribuitors"`
	// The type of the package
	Type PackageJSONTypeField `json:"type"`
	// If true, the package is considered private and will not be published
	Private bool `json:"private"`
	// An SPDX identifier of the license used by the package
	License string `json:"license"`
	// A value compared during install with process.platform
	OS []string `json:"os"`
	// A value compared during install with process.arch
	CPU []string `json:"cpu"`
	// A value compared during install with the host standard C library
	Libc string `json:"libc"`
	// The path that will be used to resolve the qualified path to use when accessing the package by its name
	Main string `json:"main"`
	// The path that will be used when an ES6-compatible environment will try to access the package by its name
	Module string `json:"module"`
	// A field used to expose some executable Javascript files to the parent package. Any entry listed here will be made available through the $PATH
	Bin map[string]string `json:"bin"`
	// A field used to list small shell scripts that will be executed when running dnpm run
	Scripts map[string]string `json:"scripts"`
	// The set of dependencies that must be made available to the current package in order for it to work properly
	Dependencies map[string]string `json:"dependencies"`
	// This field is usually not what you're looking for, unless you depend on the fsevents package. If you need a package to be required only when a specific feature is used then use an optional peer dependency
	OptionalDependencies map[string]string `json:"optionalDependencies"`
	// Similar to the dependencies field, except that these dependencies are only installed on local installs and will never be installed by the consumers of your package
	DevDependencies map[string]string `json:"devDependencies"`
	// Peer dependencies are inherited dependencies - the consumer of your package will be tasked to provide them
	PeerDependencies map[string]string `json:"peerDependencies"`
	// Workspaces are an optional feature used by monorepos to split a large project into semi-independent subprojects, each one listing their own set of dependencies
	Workspaces []string `json:"workspaces"`
	// The optional files field is an array of file patterns that describes the entries to be included when your package is installed as a dependency
	Files []string `json:"files"`
	// Specify the place where your code lives. This is helpful for people who want to contribute
	Repository PackageJSONRepositoryField `json:"repository"`
}
