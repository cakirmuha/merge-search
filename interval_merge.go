package merge_search

import (
	"sort"
)

type Interval struct {
	Start int64
	End   int64
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MergeNInterval(list ...Interval) []Interval {
	// Sort by start time descending
	sort.Slice(list, func(i, j int) bool {
		return list[i].Start > list[j].Start
	})

	pnt := 0
	for _, l := range list {
		ok := false
		for pnt > 0 && list[pnt-1].Start <= l.End {
			ok = true
			pnt--
			list[pnt].Start = min(list[pnt].Start, l.Start)
			list[pnt].End = max(list[pnt].End, l.End)
		}
		if !ok {
			list[pnt] = l
		}

		pnt++
	}

	// do not extra memory for the output
	return list[:pnt]
}
