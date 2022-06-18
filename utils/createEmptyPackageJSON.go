package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/davipatricio/colors/colors"
)

type PackageJSONFormat struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Version         string            `json:"version"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func CreateEmptyPackageJSON() error {
	// Create the package.json file
	packageJSONFile, err := os.Create("./package.json")
	if err != nil {
		return err
	}

	var buf *bytes.Buffer
	enc := json.NewEncoder(buf)
	// json.MarshalIdent escapes characters like <, > and &
	// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	err = enc.Encode(PackageJSONFormat{
		Name:        "dnpm-project",
		Description: "A project created using dnpm.",
		Version:     "0.0.1",
		Author:      "",
		License:     "MIT",
		Scripts: map[string]string{
			"test":  `echo "Error: no test specified" && exit 1`,
			"start": "node index.js",
		},
		Dependencies:    map[string]string{},
		DevDependencies: map[string]string{},
	})
	if err != nil {
		panic(err)
	}

	// Write the package.json file
	_, err = packageJSONFile.Write(buf.Bytes())
	if err != nil {
		CouldNotCreateEmptyPkgJson()
		return err
	}

	err = packageJSONFile.Close()
	if err != nil {
		panic(err)
	}

	return err
}

func CouldNotCreateEmptyPkgJson() {
	if ShowEmojis() {
		couldNotWriteEmptyPkgJsonRaw()
	} else {
		couldNotWriteEmptyPkgJsonEmojis()
	}
}

func couldNotWriteEmptyPkgJsonRaw() {
	fmt.Println(colors.Red("Could not write package.json file."))
}

func couldNotWriteEmptyPkgJsonEmojis() {
	fmt.Println("❌ ", colors.Red("Could not write package.json file."))
}
