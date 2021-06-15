package main

import (
	"io/ioutil"
	"log"
	"net/http"

	brain "trigrams/packages/learn"
)

var m *brain.Memory

func init() {
	// create memory as global variable at the beginning of the program
	m = brain.MakeMemory()
}

func main() {

	http.HandleFunc("/learn", learn)
	http.HandleFunc("/generate", generate)

	log.Println("\n\nStarting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func learn(w http.ResponseWriter, r *http.Request) {

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	mime := r.Header.Get("Content-Type")

	if mime != "text/plain" {
		log.Printf("Invalid Content-Type: %s", mime)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := m.Learn(string(resp)); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
	if err := m.Generate(w); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
