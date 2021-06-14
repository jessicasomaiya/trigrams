package learn

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSuffix(t *testing.T) {
	for _, tc := range []struct {
		name     string
		trigram  trigram
		expected string
	}{
		{
			name:     "happy",
			trigram:  "to be, that",
			expected: "be, that",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			tr := tc.trigram
			assert.Equal(t, tc.expected, tr.getSuffix())

		})
	}

}

func TestLearn(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected map[string][]trigram
	}{
		{
			name:  "without punctuation",
			input: "to be or not to be that is the question",
			expected: map[string][]trigram{
				"to be":   {"to be or", "to be that"},
				"be or":   {"be or not"},
				"or not":  {"or not to"},
				"not to":  {"not to be"},
				"be that": {"be that is"},
				"that is": {"that is the"},
				"is the":  {"is the question"},
			},
		},
		{
			name:  "with punctuation",
			input: "to be or not to be, that is the question",
			expected: map[string][]trigram{
				"to be":    {"to be or"},
				"be or":    {"be or not"},
				"or not":   {"or not to"},
				"not to":   {"not to be,"},
				"to be,":   {"to be, that"},
				"be, that": {"be, that is"},
				"that is":  {"that is the"},
				"is the":   {"is the question"},
			},
		},
		{
			name:  "with new line",
			input: "to be or not to be, that is\nthe question",
			expected: map[string][]trigram{
				"to be":    {"to be or"},
				"be or":    {"be or not"},
				"or not":   {"or not to"},
				"not to":   {"not to be,"},
				"to be,":   {"to be, that"},
				"be, that": {"be, that is"},
				"that is":  {"that is the"},
				"is the":   {"is the question"},
			},
		},
		{
			name:  "with spaces",
			input: "to be    or not to be, that is\n\nthe       question",
			expected: map[string][]trigram{
				"to be":    {"to be or"},
				"be or":    {"be or not"},
				"or not":   {"or not to"},
				"not to":   {"not to be,"},
				"to be,":   {"to be, that"},
				"be, that": {"be, that is"},
				"that is":  {"that is the"},
				"is the":   {"is the question"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := MakeMemory()
			m.Learn(tc.input)

			assert.Equal(t, tc.expected, m.brain)
		})
	}
}

func TestRun(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		prefix   string
		expected []string
	}{
		{
			name:   "More than one option",
			input:  "to be or not to be, that is the question",
			prefix: "to be",
			expected: []string{
				"to be or not to be, that is the question ",
				"to be that is the question ",
			},
		},
		{
			name:   "with newline",
			input:  "to be or not to be, that is\nthe question",
			prefix: "to be",
			expected: []string{
				"to be or not to be, that is the question ",
				"to be that is the question ",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			m := MakeMemory()

			m.Learn(tc.input)
			m.Run("to be", &buf)

			exp := tc.expected[0]
			if len(buf.String()) == 26 {
				exp = tc.expected[1]
			}

			assert.Equal(t, exp, buf.String())
		})
	}

}
