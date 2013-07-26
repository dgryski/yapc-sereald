package main

import (
	"encoding/base64"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		blob64 := r.PostFormValue("sereal")
		b, _ := base64.StdEncoding.DecodeString(blob64)
		sdecode := sereal.Decoder{}
		var intf interface{}
		sdecode.Unmarshal(b, &intf)
		w.Header().Set("Content-Type", "text/plain")
		spew.Fdump(w, intf)
	})

	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	log.Fatal(http.ListenAndServe(port, nil))
}
