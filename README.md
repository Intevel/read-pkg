# read-pkg

🌠 Rapid fast reading of package.json files in Go

## Installation

```bash
go get github.com/intevel/read-pkg
```

## Example

```go
package main

import (
	"log"
	p "github.com/intevel/read-pkg"
)

func main () {
	pkg, err := p.ParsePackage()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pkg)
}
```
