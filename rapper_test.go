package wackrapper

import (
	"fmt"
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	rappers := make([]Rapper, 10)
	for i := 0; i < 10; i++ {
		rappers[i] = New()
	}
	// TODO: Add better test
}

func TestWackness(t *testing.T) {
	for _, tc := range testWackRappers {
		w := tc.rapper.Wackness()
		fmt.Printf("%v's wackness: %v\n", tc.rapper.name, w)
		if math.IsNaN(w) {
			t.Errorf("FAIL %v - got %v want any number", tc.description, w)
		}
	}
}
