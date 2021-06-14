package learn

// add error handing and concurrency

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Memory holds a map in whcih the key is a string of 2 words and the value
//  is a slice of trigrams that have the key as a prefix.
type Memory struct {
	brain map[string][]trigram // holds all thre trigarams
	rand  *rand.Rand
}

func MakeMemory() *Memory {
	return &Memory{
		brain: make(map[string][]trigram),
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// a trigram is are strings with 3 words
type trigram string

// write writes the last word in the trigram
func (tr trigram) write(w io.Writer) {
	fmt.Fprintf(w, "%s ",
		strings.Split(string(tr), " ")[2],
	)
}

// Learn populates the brain in memory which is a map. The key is a string of 2 words and the value
//  is a slice of trigrams that have the key as a prefix.
func (m *Memory) Learn(input string) error {

	if input == "" {
		return errors.New("empty input - please provide an input")
	}

	// remove new lines and spaces
	reg := regexp.MustCompile(`\s+`)
	input = reg.ReplaceAllString(input, " ")

	inputSlice := strings.Split(input, " ")
	for k := range inputSlice {

		// cannot make trigram with < 3 words because you cannot make more trigrams
		if k == len(inputSlice)-2 {
			return nil
		}

		// build trigram using slice of strings
		var trig []string
		for i := 0; i < 3; i++ {
			trig = append(trig, inputSlice[k+i])
		}

		// create key using first 2 words of trigram
		key := strings.Join([]string{trig[0], trig[1]}, " ")

		m.brain[key] = append(m.brain[key], trigram(strings.Join(trig, " ")))
	}

	return nil
}

func (m *Memory) Generate(w io.Writer) {
	for k := range m.brain {
		//  k is random key
		m.Run(k, w)
		break
	}
}

// Run generates a block of text using the brain in Memory. Run takes the prefix as a starting point,
// and writes the block of text to w until no trigrams exist at given key.
func (m *Memory) Run(prefix string, w io.Writer) {

	// write initial strings before algo begins
	fmt.Fprintf(w, "%s ", prefix)

	for {

		// get slice of trigrams with given prefix - if no trigram exists then end the algorithm
		trigs, ok := m.brain[prefix]
		if !ok || len(trigs) == 0 {
			return
		}

		// select new trigram out of trigs by using a random index
		idx := m.rand.Intn(len(trigs))
		newTrigram := trigs[idx]

		newTrigram.write(w)

		// delete selected to avoid repeats
		m.brain[prefix] = append(m.brain[prefix][:idx], m.brain[prefix][idx+1:]...)

		// use the last 2 words of the new trigram as the new prefix
		prefix = newTrigram.getSuffix()
	}

}

// getSuffix returns the last 2 words in a trigram
func (tr trigram) getSuffix() string {
	return strings.Join(
		strings.Split(string(tr), " ")[1:3],
		" ",
	)
}
