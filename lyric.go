package wackrapper

type lyric struct {
	id       int
	song     string
	verse    verse
	wackness float64
}

type verse struct {
	line []string
}

// TODO: youtube link in lyric struct
