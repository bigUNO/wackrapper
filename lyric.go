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

// Judge accepts a 1 or -1
func (l *lyric) Judge(v int) {
	switch v {
	case -1:
		l.wackness--
	case 1:
		l.wackness++
	default:
		return
	}
}

// TODO: youtube link in lyric struct
