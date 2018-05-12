// Package wackrapper is a inside joke.
// Creates rappers and links them to their wack lines.
package wackrapper

import (
	"fmt"
	"math"
	"sort"

	"github.com/rs/xid"
)

// Rapper is the sum of their wack rhymes
type Rapper struct {
	xid      xid.ID
	name     string
	wackness float64
	lyrics   []lyric
}

func (r Rapper) Wackness() float64 {
	var total float64 = 0
	for _, v := range r.lyrics {
		total += v.wackness
	}
	wack := total / float64(len(r.lyrics))
	if math.IsNaN(wack) {
		return 0
	}
	return wack
}

// New returns a new rapper
func New() Rapper {
	return Rapper{xid: xid.New()}
}

// Rank returns an ordered array of rappers in order
// of wackness. The first rapper is the wackest in
// the group.
func Rank(rr []Rapper) []string {

	type ranking struct {
		name     string
		wackness float64
	}

	switch len(rr) {
	case 0:
		return []string{}
	default:
		m := make(map[int]Rapper)
		// Load map with rappers
		for k, v := range rr {
			m[k] = v
		}

		// Load map keys into a slice
		keys := []float64{}
		for i := range m {
			keys = append(keys, m[i].Wackness())
		}
		// Sort slice
		//sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		//orderedKeys := make([]float64, len(keys))
		//for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		//orderedKeys[i], orderedKeys[j] = keys[j], keys[i]
		//}
		for i := len(keys)/2 - 1; i >= 0; i-- {
			opp := len(keys) - 1 - i
			keys[i], keys[opp] = keys[opp], keys[i]
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		fmt.Printf("%v\n", keys)
		// Build slice of strings based on map now that
		// values are in order
		ss := []string{}
		for k := range keys {
			ss = append(ss, m[k].name)
		}
		fmt.Printf("%v\n", ss)
		/*
			rr := []ranking{}
			for k := range keys {
				rr = append(rr, ranking{m[k].name, m[k].Wackness()})
			}

			ss := []string{}
			for _, v := range rr {
				//ss = append(ss, v.name, strconv.FormatFloat(v.wackness, 'f', 2, 64))
				ss = append(ss, v.name)
			}*/

		return ss
	}

}
