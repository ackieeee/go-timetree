package main

import (
	"fmt"
	"github.com/gba-3/go-timetree/timetree"
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("TIMETREE_AUTH_TOKEN")
	cli := timetree.NewClient(&http.Client{}, token)
	data, err := cli.Calendar.List([]string{"labels", "members"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", data)
}
