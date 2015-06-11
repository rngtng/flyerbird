package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"os/exec"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()

	m.Use(martini.Static("public"))

	m.Post("/render", func(response http.ResponseWriter, req *http.Request) {
		dataFront := req.FormValue("data_front")
		data, err := base64.StdEncoding.DecodeString(dataFront)
		if err != nil {
			log.Fatal(err)
		}
		fileFront, _ := ioutil.TempFile(os.TempDir(), "front")
		defer os.Remove(fileFront.Name())
		fileFront.Write(data)

		dataBack := req.FormValue("data_back")
		data, err = base64.StdEncoding.DecodeString(dataBack)
		if err != nil {
			log.Fatal(err)
		}
		fileBack, _ := ioutil.TempFile(os.TempDir(), "back")
		defer os.Remove(fileBack.Name())
		fileBack.Write(data)

		out, err := exec.Command("cut/main.py", fileFront.Name(), fileBack.Name()).CombinedOutput()
		if err != nil {
			log.Println(string(out))
			log.Fatal(err)
		}

		response.Header().Set("Content-Type", "application/octet-stream")
		response.Header().Set("Content-Disposition", "attachment; filename=\"origami.pdf\"")
		response.Write(out)
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
