package wackrapper

import "testing"

func TestNew(t *testing.T) {
	rappers := make([]Rapper, 10)
	for i := 0; i < 10; i++ {
		rappers[i] = New()
	}
	// TODO: Add better test
}
