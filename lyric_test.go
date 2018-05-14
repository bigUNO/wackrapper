package wackrapper

import "testing"

func TestJudge(t *testing.T) {
	for _, tc := range testLyrics {

		tc.lyric.Judge(int(tc.judge[0]))
		if int(tc.lyric.wackness) != tc.judge[1] {
			t.Errorf("FAIL judging lyrics from %v broke got %v want %v", tc.lyric.song, tc.lyric.wackness, tc.judge[1])
		}
	}
}
