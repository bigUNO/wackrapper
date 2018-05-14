package wackrapper

import "github.com/rs/xid"

var id = xid.ID{0x4d, 0x88, 0xe1, 0x5b, 0x60, 0xf4, 0x86, 0xe4, 0x28, 0x41, 0x2d, 0xc9}

var testLyrics = []struct {
	description string
	lyric       lyric
	judge       []int
}{
	{
		"invisible lyric",
		lyric{
			1,
			"The song that doesn't exist",
			verse{
				[]string{""},
			},
			0,
		},
		[]int{666, 0},
	},
	{
		"single-verse lyric",
		lyric{
			1,
			"Pop Style",
			verse{
				[]string{"Got so many chains they call me Chaining Tatum"},
			},
			0,
		},
		[]int{1, 1},
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
		[]int{-1, 22},
	},
}

var testWackRappers = []struct {
	description string
	rapper      Rapper
}{
	{
		"the invisible rapper, no lyrics",
		Rapper{
			id,
			"Lil Invisible",
			0,
			[]lyric{},
		},
	},
	{
		"Rapper with one wack lyric consisting of one verse",
		Rapper{
			id,
			"Chuggo",
			0,
			[]lyric{
				{
					1,
					"Ah C'mon",
					verse{
						[]string{
							"AH! COM'ON FUCKING GUY!!",
						},
					},
					10,
				},
			},
		},
	},
	{
		"Rapper with one wack lyric consisting of two verses",
		Rapper{
			id,
			"Nicki Minaj",
			0,
			[]lyric{
				{
					1,
					"Your Love",
					verse{
						[]string{
							"When I was a geisha he was a samuari",
							"Somehow I understood him when he spoke Thai",
						},
					},
					16,
				},
			},
		},
	},
	{
		"Rapper with multiple wack lyrics",
		Rapper{
			id,
			"Canibus",
			0,
			[]lyric{
				{
					1,
					"Funk Master Flex Freestyle",
					verse{
						[]string{
							"I can double my density from three-sixty degrees to seven-twenty instantly",
						},
					},
					16,
				},
				{
					1,
					"Second Round KO",
					verse{
						[]string{
							"You might got more cash than me",
							"but you ain't got the skills to eat a nigga's ass like me",
						},
					},
					1567,
				},
			},
		},
	},
}
