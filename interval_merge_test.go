package merge_search

import (
	"testing"
)

func TestMergeNInterval(t *testing.T) {
	tests := []struct {
		input  []Interval
		output []Interval
	}{
		{
			input:  []Interval{{1, 5}, {2, 4}, {5, 8}},
			output: []Interval{{1, 8}},
		},
		{
			input:  []Interval{{1, 3}, {4, 9}, {5, 10}},
			output: []Interval{{4, 10}, {1, 3}},
		},
		{
			input:  []Interval{{5, 7}, {1, 4}, {8, 11}, {3, 5}},
			output: []Interval{{8, 11}, {1, 7}},
		},
	}

	for i, test := range tests {
		got := MergeNInterval(test.input...)
		if want, have := len(test.output), len(got); want != have {
			t.Fatalf("#%v: expected %v elements; got: %v", i, want, have)
		}
		for i := 0; i < len(got); i++ {
			if want, have := test.output[i], got[i]; want != have {
				t.Errorf("#%v: expected %v; got: %v", i, want, have)
			}
		}
	}
}
