package readPackage

import (
	"encoding/json"
	"io/ioutil"
    "log"
	"os"
)

var requiredFields = []string{"name", "version"}

type Contributor struct {
    Name string `json:"name"`
	Email string `json:"email"`
	Url string `json:"url"`
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Private bool `json:"private"`
	Description string `json:"description"`
	Keywords []string `json:"keywords"`
	Homepage string `json:"homepage"`
	Author string `json:"author"`
	Contributors []Contributor `json:"contributors"`
	License string `json:"license"`
	Repository string `json:"repository"`
	Dependencies map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	PeerDependencies map[string]string `json:"peerDependencies"`
}

func parsePackage() Package {
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

	return data
}