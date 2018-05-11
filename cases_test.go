package wackrapper

import "github.com/rs/xid"

var id = xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}

var testLyrics = []struct {
	description string
	lyric       lyric
}{
	{
		"single-verse lyric",
		lyric{
			1,
			"Pop Style",
			verse{
				[]string{"Got so many chains they call me Chaining Tatum"},
				//},
			},
			0,
		},
	},
	{
		"multi-verse lyrics",
		lyric{
			2,
			"Your Love",
			verse{
				[]string{
					"When I was a geisha he was a samuari",
					"Somehow I understood him when he spoke Thai",
				},
			},
			23,
		},
	},
}

/*
var testWackRappers = []struct {
	description string
	rapper      Rapper
}{
	{
		description: "destfr",
		Rapper{
			xid:      id,
			name:     "Chuggo",
			wackness: 0,
			[]lyrics{
				id:   1,
				song: "Ah C'mon",
				verse[line:"My brother's in the can and won't get out till next July."],
				wackness: 3.4,
			},
		},
	},
}
*/
