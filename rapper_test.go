package wackrapper

import (
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

func benchmarkNew(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if n < i {
			New()
		}
	}
}

func BenchmarkNew10(b *testing.B)         { benchmarkNew(10, b) }
func BenchmarkNew1000(b *testing.B)       { benchmarkNew(1000, b) }
func BenchmarkNew1000000(b *testing.B)    { benchmarkNew(1000000, b) }
func BenchmarkNew1000000000(b *testing.B) { benchmarkNew(1000000000, b) }

func TestWackness(t *testing.T) {
	for _, tc := range testWackRappers {
		w := tc.rapper.Wackness()
		//fmt.Printf("%v's wackness: %v\n", tc.rapper.name, w)
		if math.IsNaN(w) {
			t.Errorf("FAIL %v - got %v want any number", tc.description, w)
		}
	}
}

func TestRank(t *testing.T) {
	want := []string{"Canibus", "Nicki Minaj", "Chuggo", "Lil Invisible"}
	rr := []Rapper{}
	//for _, tc := range testWackRappers {
	//rr :=
	//}
	empty := Rank(rr)
	if len(empty) != 0 {
		t.Errorf("FAIL wrong lenth list; want 0 got %v", len(empty))
	}

	for _, tc := range testWackRappers {
		rr = append(rr, tc.rapper)
	}
	list := Rank(rr)
	switch {
	case len(list) != len(want):
		t.Errorf("FAIL got wrong lenth Rank() response want %v got %v", len(want), len(list))
	case len(list) == len(want):
		for i := range list {
			if list[i] != want[i] {
				t.Errorf("FAIL got wrong order; want %v got %v", want[i], list[i])
			}
		}
	}

}
