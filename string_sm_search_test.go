package merge_search

import (
	"testing"
)

// Test if the given string can be represented as a concatenation of any subset of tokens
func TestCheckSingleString(t *testing.T) {
	tests := []struct {
		tokens []string
		input  string
		output bool
	}{
		{
			tokens: []string{"ab", "bc", "a"},
			input:  "abc",
			output: true,
		},
		{
			tokens: []string{"a", "ab", "bc"},
			input:  "ab",
			output: true,
		},
		{
			tokens: []string{"a", "ab", "bc"},
			input:  "",
			output: true,
		},
		{
			tokens: []string{"ab", "bc"},
			input:  "abc",
			output: false,
		},
		{
			tokens: []string{"ab", "bc", "c"},
			input:  "abbcbc",
			output: true,
		},
		{
			tokens: []string{"ab", "bc", "cd", "de"},
			input:  "abbcbcde",
			output: true,
		},
		{
			tokens: []string{"ab", "bc", "cd", "de", "ec"},
			input:  "abbcbccdecde",
			output: true,
		},
		{
			tokens: []string{"ab", "bc", "cd", "de", "ec"},
			input:  "abbcbccdeccde",
			output: false,
		},
	}

	for i, test := range tests {
		got := CheckSingleString(test.tokens, test.input)
		if want, have := test.output, got; want != have {
			t.Fatalf("#%v: expected %v elements; got: %v", i, want, have)
		}
	}
}

// Test if long input strings can be represented as a concatenation of any subset of tokens
/*func TestMultipleStringSearch(t *testing.T) {
	tests := []struct {
		tokens     []string
		inputList  []string
		outputList []bool
	}{
		{
			tokens:     []string{"ab", "bc", "aa"},
			inputList:  []string{"ab", "abc", "abaaaabc", "abbcbc"},
			outputList: []bool{true, false, true, true},
		},
		{
			tokens:     []string{"b", "bcc", "aba"},
			inputList:  []string{"ababccbaba", "bbccb", "ababbbcc", "abaab", "abaababbcc"},
			outputList: []bool{true, true, true, false, true},
		},
	}

	for i, test := range tests {
		got := CheckInList(test.tokens, test.inputList)
		if want, have := len(test.outputList), len(got); want != have {
			t.Fatalf("#%v: expected %v elements; got: %v", i, want, have)
		}
		for i := 0; i < len(got); i++ {
			if want, have := test.outputList[i], got[i]; want != have {
				t.Errorf("#%v: expected %v; got: %v", i, want, have)
			}
		}
	}
}*/
