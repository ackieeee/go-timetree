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
	cli := timetree.NewClient(&http.Client{}, token)
	data, err := cli.Calendar.List([]string{"labels", "members"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", data)
}
