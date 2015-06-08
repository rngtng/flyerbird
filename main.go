package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
	"os/exec"
	// "strings"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()

	m.Get("/render", func(req *http.Request) []byte {
		// params := req.URL.Query()
		// urls := strings.Split(params.Get("url"), "\n")

		command := "cat"
		params := "data/test.pdf"

		out, err := exec.Command(command, params).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		return out
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
