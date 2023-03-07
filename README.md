# go-timetree
[![Test](https://github.com/sugartr3e/go-timetree/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gba-3/go-timetree/actions/workflows/test.yml)

Go client library for Timetree API

## Requirements
- Go >= 1.19

## Installation
```
go get github.com/sugartr3e/go-timetree
```

Examples
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sugartr3e/go-timetree/timetree"
)

func main() {
	token := os.Getenv("TIMETREE_AUTH_TOKEN")
	if token == "" {
		log.Fatal("timetree token is empty")
	}
	cli := timetree.NewClient(&http.Client{}, token)
	data, err := cli.Calendar.List([]string{"labels", "members"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", data)
}
```
