package main

import (
	"fmt"
	"github.com/gba-3/go-timetree/timetree"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("TIMETREE_AUTH_TOKEN")
	cli := timetree.NewClient(&http.Client{}, token)
	fmt.Println(cli.BaseURL)
}
