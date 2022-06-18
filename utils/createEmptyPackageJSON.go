package utils

import (
	"bytes"
	"dnpm/structs"
	"encoding/json"
	"fmt"
	"os"

	"github.com/davipatricio/colors/colors"
)

func CreateEmptyPackageJSON() error {
	// Create the package.json file
	packageJSONFile, err := os.Create("./package.json")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	// json.MarshalIdent escapes characters like <, > and &
	// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	err = enc.Encode(structs.PackageJSONFormat{
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
