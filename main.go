package main

import (
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()

	m.Get("/", func(response http.ResponseWriter) {
		response.Write([]byte("Hello World"))
	})

	m.Get("/render", func(response http.ResponseWriter, req *http.Request) {
		// params := req.URL.Query()
		// urls := strings.Split(params.Get("url"), "\n")

		command := "cat"
		params := "data/test.pdf"

		out, err := exec.Command(command, params).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		response.Header().Set("Content-Type", "application/octet-stream")
		response.Write(out)
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
