package mrctech

import (
	"fmt"
	"sort"
)

type P struct {
	In  int
	Out int
}

// Assume P is validated
func DurationPeak(ps []P, durationIn int, durationOut int, display bool) int {
	sub := durationOut - durationIn
	pm := make([]int, sub+1)
	for _, p := range ps {
		recordBegin := p.In
		recordEnd := p.Out
		if recordBegin > durationOut || recordEnd < durationIn {
			continue
		}
		if recordBegin < durationIn {
			recordBegin = durationIn
		}
		if recordEnd > durationOut {
			recordEnd = durationOut
			pm[sub] = pm[sub] + 1
		}
		for i := recordBegin - durationIn; i+durationIn < recordEnd; i++ {
			pm[i] = pm[i] + 1
		}
	}

	if display {
		fmt.Println(pm)
	}
	sort.Ints(pm)

	return pm[sub]
}
