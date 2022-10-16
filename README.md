# go-timetree
[![codecov](https://codecov.io/gh/gba-3/go-timetree/branch/main/graph/badge.svg?token=OCMC7KZJ4I)](https://codecov.io/gh/gba-3/go-timetree)
[![Test](https://github.com/gba-3/go-timetree/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gba-3/go-timetree/actions/workflows/test.yml)

Go client library for Timetree API

## Requirements
- Go >= 1.19

## Installation
```
go install github.com/gba-3/go-timetree
```

Examples
```go
package main

import (
	"fmt"
	"log"
	"github.com/gba-3/go-timetree/timetree"
	"net/http"
	"os"
)

func main() {
	// The token for request to Timetree API.
	token := os.Getenv("TIMETREE_AUTH_TOKEN")
	
	// Create Timetree API Client.
	cli := timetree.NewClient(&http.Client{}, token)
	r, err := cli.Calendar.List([]string{"labels", "members"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", r)
}
```
