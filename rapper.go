// Package wackrapper is a inside joke.
// Creates rappers and links them to their wack lines.
package wackrapper

import (
	"math"

	"github.com/rs/xid"
)

// Rapper is the sum of their wack rhymes
type Rapper struct {
	xid      xid.ID
	name     string
	wackness float64
	lyrics   []lyric
}

func (r *Rapper) Wackness() float64 {
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
