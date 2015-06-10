package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-martini/martini"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))

	m.Get("/", func(response http.ResponseWriter) {
		response.Write([]byte("Hello World"))
	})

	m.Post("/render", func(response http.ResponseWriter, req *http.Request) {
		dataFront := req.FormValue("data_front")
		data, err := base64.StdEncoding.DecodeString(dataFront)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("front.png", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// file, _ := ioutil.TempFile(os.TempDir(), "front")
		// defer os.Remove(file.Name())
		// file.Write(data)

		dataBack := req.FormValue("data_back")
		data, err = base64.StdEncoding.DecodeString(dataBack)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("back.png", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// file, _ := ioutil.TempFile(os.TempDir(), "back")
		// defer os.Remove(file.Name())
		// file.Write(data)

		command := "cat"
		cliParams := "data/test.pdf"

		out, err := exec.Command(command, cliParams).CombinedOutput()
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
