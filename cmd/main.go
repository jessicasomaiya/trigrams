package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	brain "trigrams/packages/learn"
)

var m *brain.Memory

func init() {
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
		http.Error(w, err.Error(), http.StatusBadRequest) //change
		return
	}

	input := string(resp)
	if err := m.Learn(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //change
		return
	}

	fmt.Fprintf(w, "Thank you, lesson learned")
}

func generate(w http.ResponseWriter, r *http.Request) {
	m.Generate(w)
}
