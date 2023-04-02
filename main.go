package main

import (
	"encoding/json"
	"io/ioutil"
    "log"
	"os"
)

var requiredFields = []string{"name", "version"}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Private bool `json:"private"`
	Description string `json:"description"`
	Keywords []string `json:"keywords"`
	Homepage string `json:"homepage"`
	Author any `json:"author"`
	Contributors any `json:"contributors"`
	Funding any `json:"funding"`
	License string `json:"license"`
	Main string `json:"main"`
	Bin any `json:"bin"`
	Browser any `json:"browser"`
	Man string `json:"man"`
	Directories any `json:"directories"`
	Files       []string          `json:"files"`
	Scripts     map[string]string `json:"scripts"`
	Os          []string          `json:"os"`
	Cpu         []string          `json:"cpu"`
	Repository string `json:"repository"`
	Dependencies map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	PeerDependencies map[string]string `json:"peerDependencies"`
	BundleDependencies map[string]string `json:"bundleDependencies"`
	OptionalDependencies map[string]string `json:"optionalDependencies"`
}

func ParsePackage() (Package, error) {
	file, err := os.Open("package.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    var data Package
    err = json.Unmarshal(byteValue, &data)
    if err != nil {
        log.Fatal(err)
    }

	return data, err
}