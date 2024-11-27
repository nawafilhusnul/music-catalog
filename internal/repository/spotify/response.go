package spotify

var searchResponse = `
{
  "albums": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=album&market=ID&locale=en-US,en;q%3D0.7",
    "limit": 10,
    "next": "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=album&market=ID&locale=en-US,en;q%3D0.7",
    "offset": 0,
    "previous": null,
    "total": 897,
    "items": [
      {
        "album_type": "album",
        "total_tracks": 22,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE"
        },
        "href": "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
        "id": "6i6folBtxKV28WX3msQ4FE",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody (The Original Soundtrack)",
        "release_date": "2018-10-19",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:6i6folBtxKV28WX3msQ4FE",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/1ROGEes7F9TOfgJWZ5rBZC"
        },
        "href": "https://api.spotify.com/v1/albums/1ROGEes7F9TOfgJWZ5rBZC",
        "id": "1ROGEes7F9TOfgJWZ5rBZC",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b273ef6a9f30737924cc52faf39f",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e02ef6a9f30737924cc52faf39f",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d00004851ef6a9f30737924cc52faf39f",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody",
        "release_date": "2023-07-13",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:1ROGEes7F9TOfgJWZ5rBZC",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/48bKH1ugFBhERC1rdojP9d"
            },
            "href": "https://api.spotify.com/v1/artists/48bKH1ugFBhERC1rdojP9d",
            "id": "48bKH1ugFBhERC1rdojP9d",
            "name": "Dewa 19",
            "type": "artist",
            "uri": "spotify:artist:48bKH1ugFBhERC1rdojP9d"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/6FmAhY7J9RleYQt4cZzyCH"
        },
        "href": "https://api.spotify.com/v1/albums/6FmAhY7J9RleYQt4cZzyCH",
        "id": "6FmAhY7J9RleYQt4cZzyCH",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b27378e0a900d4800d86d6849f86",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e0278e0a900d4800d86d6849f86",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d0000485178e0a900d4800d86d6849f86",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody",
        "release_date": "2023-03-20",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:6FmAhY7J9RleYQt4cZzyCH",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/2TSeIynP2u22bqZOgKkbZm"
            },
            "href": "https://api.spotify.com/v1/artists/2TSeIynP2u22bqZOgKkbZm",
            "id": "2TSeIynP2u22bqZOgKkbZm",
            "name": "Ahmad Dhani",
            "type": "artist",
            "uri": "spotify:artist:2TSeIynP2u22bqZOgKkbZm"
          },
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/4nUgYF7wmkUux1M9SdXM0h"
            },
            "href": "https://api.spotify.com/v1/artists/4nUgYF7wmkUux1M9SdXM0h",
            "id": "4nUgYF7wmkUux1M9SdXM0h",
            "name": "Philharmonic Orchestra",
            "type": "artist",
            "uri": "spotify:artist:4nUgYF7wmkUux1M9SdXM0h"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/66gjGuZX4OGM8CrlfIMftl"
        },
        "href": "https://api.spotify.com/v1/albums/66gjGuZX4OGM8CrlfIMftl",
        "id": "66gjGuZX4OGM8CrlfIMftl",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b273aed306b9bad470242efb68b8",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e02aed306b9bad470242efb68b8",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d00004851aed306b9bad470242efb68b8",
            "width": 64
          }
        ],
        "name": "bohemian rhapsody - guitar version",
        "release_date": "2022-11-30",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:66gjGuZX4OGM8CrlfIMftl",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/5Uspdg4EAuwEHmjPtQD8NP"
            },
            "href": "https://api.spotify.com/v1/artists/5Uspdg4EAuwEHmjPtQD8NP",
            "id": "5Uspdg4EAuwEHmjPtQD8NP",
            "name": "guitar girl",
            "type": "artist",
            "uri": "spotify:artist:5Uspdg4EAuwEHmjPtQD8NP"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/58seMZR64MQxmQHH20LRUQ"
        },
        "href": "https://api.spotify.com/v1/albums/58seMZR64MQxmQHH20LRUQ",
        "id": "58seMZR64MQxmQHH20LRUQ",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b2738891c122e4f79f262e21ebd5",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e028891c122e4f79f262e21ebd5",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d000048518891c122e4f79f262e21ebd5",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody",
        "release_date": "2020-01-24",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:58seMZR64MQxmQHH20LRUQ",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/0iQDOaYEA5i9RAF0Z73iXb"
            },
            "href": "https://api.spotify.com/v1/artists/0iQDOaYEA5i9RAF0Z73iXb",
            "id": "0iQDOaYEA5i9RAF0Z73iXb",
            "name": "Angelina Jordan",
            "type": "artist",
            "uri": "spotify:artist:0iQDOaYEA5i9RAF0Z73iXb"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/6mvwcq1i7w481wYCtxJyso"
        },
        "href": "https://api.spotify.com/v1/albums/6mvwcq1i7w481wYCtxJyso",
        "id": "6mvwcq1i7w481wYCtxJyso",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b273dbf9bbde43127734e1ce9c29",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e02dbf9bbde43127734e1ce9c29",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d00004851dbf9bbde43127734e1ce9c29",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody (Acoustic)",
        "release_date": "2018-10-22",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:6mvwcq1i7w481wYCtxJyso",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/4LaimQU44rsz2kMWQmY6Bi"
            },
            "href": "https://api.spotify.com/v1/artists/4LaimQU44rsz2kMWQmY6Bi",
            "id": "4LaimQU44rsz2kMWQmY6Bi",
            "name": "John Adams",
            "type": "artist",
            "uri": "spotify:artist:4LaimQU44rsz2kMWQmY6Bi"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/73Td5oIMISZDGAvzGvEpbC"
        },
        "href": "https://api.spotify.com/v1/albums/73Td5oIMISZDGAvzGvEpbC",
        "id": "73Td5oIMISZDGAvzGvEpbC",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b27340cead1efe274e2658877624",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e0240cead1efe274e2658877624",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d0000485140cead1efe274e2658877624",
            "width": 64
          }
        ],
        "name": "Bohemian Rhapsody",
        "release_date": "2024-11-12",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:73Td5oIMISZDGAvzGvEpbC",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/32Jb1X3wSmmoHj2epZReZA"
            },
            "href": "https://api.spotify.com/v1/artists/32Jb1X3wSmmoHj2epZReZA",
            "id": "32Jb1X3wSmmoHj2epZReZA",
            "name": "Steve Vai",
            "type": "artist",
            "uri": "spotify:artist:32Jb1X3wSmmoHj2epZReZA"
          },
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/4IVXzZW6i9KIcpeP29Y1Df"
            },
            "href": "https://api.spotify.com/v1/artists/4IVXzZW6i9KIcpeP29Y1Df",
            "id": "4IVXzZW6i9KIcpeP29Y1Df",
            "name": "Generation Axe",
            "type": "artist",
            "uri": "spotify:artist:4IVXzZW6i9KIcpeP29Y1Df"
          },
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/2NcbLU1bW55eahD0UgD7U3"
            },
            "href": "https://api.spotify.com/v1/artists/2NcbLU1bW55eahD0UgD7U3",
            "id": "2NcbLU1bW55eahD0UgD7U3",
            "name": "Brian May",
            "type": "artist",
            "uri": "spotify:artist:2NcbLU1bW55eahD0UgD7U3"
          }
        ]
      },
      {
        "album_type": "album",
        "total_tracks": 11,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/3LRdN6TlxSk6DJqpSWwdFQ"
        },
        "href": "https://api.spotify.com/v1/albums/3LRdN6TlxSk6DJqpSWwdFQ",
        "id": "3LRdN6TlxSk6DJqpSWwdFQ",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/64825eca60656397ee6133105402d94df84af353",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/8982055798f8c8dfea1195c539cdc8204925a1e1",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/a8aa505bb39f2369af2eee428281beb3bdf82949",
            "width": 64
          }
        ],
        "name": "Bintang Lima",
        "release_date": "2000-05-15",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:3LRdN6TlxSk6DJqpSWwdFQ",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/35dWPb7Tmq0WESp6KAUJ8w"
            },
            "href": "https://api.spotify.com/v1/artists/35dWPb7Tmq0WESp6KAUJ8w",
            "id": "35dWPb7Tmq0WESp6KAUJ8w",
            "name": "Dewa",
            "type": "artist",
            "uri": "spotify:artist:35dWPb7Tmq0WESp6KAUJ8w"
          }
        ]
      },
      {
        "album_type": "single",
        "total_tracks": 1,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/5S5vnFRSXVua5SbdoY9S8y"
        },
        "href": "https://api.spotify.com/v1/albums/5S5vnFRSXVua5SbdoY9S8y",
        "id": "5S5vnFRSXVua5SbdoY9S8y",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b2730104b1f71c7bd8e6e5ea1fd2",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e020104b1f71c7bd8e6e5ea1fd2",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d000048510104b1f71c7bd8e6e5ea1fd2",
            "width": 64
          }
        ],
        "name": "Somebody's Pleasure",
        "release_date": "2023-01-27",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:5S5vnFRSXVua5SbdoY9S8y",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/6ygKuZFz2sRggPZRaLHVHD"
            },
            "href": "https://api.spotify.com/v1/artists/6ygKuZFz2sRggPZRaLHVHD",
            "id": "6ygKuZFz2sRggPZRaLHVHD",
            "name": "Aziz Hedra",
            "type": "artist",
            "uri": "spotify:artist:6ygKuZFz2sRggPZRaLHVHD"
          }
        ]
      },
      {
        "album_type": "album",
        "total_tracks": 14,
        "is_playable": true,
        "external_urls": {
          "spotify": "https://open.spotify.com/album/2hwNRtZyqmEcCzV47OUAlN"
        },
        "href": "https://api.spotify.com/v1/albums/2hwNRtZyqmEcCzV47OUAlN",
        "id": "2hwNRtZyqmEcCzV47OUAlN",
        "images": [
          {
            "height": 640,
            "url": "https://i.scdn.co/image/ab67616d0000b27369222d0908ecf5d22cc9d960",
            "width": 640
          },
          {
            "height": 300,
            "url": "https://i.scdn.co/image/ab67616d00001e0269222d0908ecf5d22cc9d960",
            "width": 300
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/ab67616d0000485169222d0908ecf5d22cc9d960",
            "width": 64
          }
        ],
        "name": "The 2000's Greatest",
        "release_date": "2007-03-12",
        "release_date_precision": "day",
        "type": "album",
        "uri": "spotify:album:2hwNRtZyqmEcCzV47OUAlN",
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/48bKH1ugFBhERC1rdojP9d"
            },
            "href": "https://api.spotify.com/v1/artists/48bKH1ugFBhERC1rdojP9d",
            "id": "48bKH1ugFBhERC1rdojP9d",
            "name": "Dewa 19",
            "type": "artist",
            "uri": "spotify:artist:48bKH1ugFBhERC1rdojP9d"
          }
        ]
      }
    ]
  }
}
`

var searchResponseStruct = &SearchResponse{
	Albums: Albums{
		HREF:  "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=album&market=ID&locale=en-US,en;q%3D0.7",
		Limit: 10,
		Next: func() *string {
			next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=album&market=ID&locale=en-US,en;q%3D0.7"
			return &next
		}(),
		Offset:   0,
		Previous: nil,
		Total:    897,
		Items: []AlbumItem{
			{
				AlbumType:   "album",
				TotalTracks: 22,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE",
				},
				Href: "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
				ID:   "6i6folBtxKV28WX3msQ4FE",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b", Width: 64},
				},
				Name:                 "Bohemian Rhapsody (The Original Soundtrack)",
				ReleaseDate:          "2018-10-19",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:6i6folBtxKV28WX3msQ4FE",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d",
						},
						Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
						ID:   "1dfeR4HaWDbWqFHLkxsg1d",
						Name: "Queen",
						Type: "artist",
						URI:  "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/1ROGEes7F9TOfgJWZ5rBZC",
				},
				Href: "https://api.spotify.com/v1/albums/1ROGEes7F9TOfgJWZ5rBZC",
				ID:   "1ROGEes7F9TOfgJWZ5rBZC",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b273ef6a9f30737924cc52faf39f", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e02ef6a9f30737924cc52faf39f", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d00004851ef6a9f30737924cc52faf39f", Width: 64},
				},
				Name:                 "Bohemian Rhapsody",
				ReleaseDate:          "2023-07-13",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:1ROGEes7F9TOfgJWZ5rBZC",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/48bKH1ugFBhERC1rdojP9d",
						},
						Href: "https://api.spotify.com/v1/artists/48bKH1ugFBhERC1rdojP9d",
						ID:   "48bKH1ugFBhERC1rdojP9d",
						Name: "Dewa 19",
						Type: "artist",
						URI:  "spotify:artist:48bKH1ugFBhERC1rdojP9d",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/6FmAhY7J9RleYQt4cZzyCH",
				},
				Href: "https://api.spotify.com/v1/albums/6FmAhY7J9RleYQt4cZzyCH",
				ID:   "6FmAhY7J9RleYQt4cZzyCH",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b27378e0a900d4800d86d6849f86", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e0278e0a900d4800d86d6849f86", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d0000485178e0a900d4800d86d6849f86", Width: 64},
				},
				Name:                 "Bohemian Rhapsody",
				ReleaseDate:          "2023-03-20",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:6FmAhY7J9RleYQt4cZzyCH",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/2TSeIynP2u22bqZOgKkbZm",
						},
						Href: "https://api.spotify.com/v1/artists/2TSeIynP2u22bqZOgKkbZm",
						ID:   "2TSeIynP2u22bqZOgKkbZm",
						Name: "Ahmad Dhani",
						Type: "artist",
						URI:  "spotify:artist:2TSeIynP2u22bqZOgKkbZm",
					},
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/4nUgYF7wmkUux1M9SdXM0h",
						},
						Href: "https://api.spotify.com/v1/artists/4nUgYF7wmkUux1M9SdXM0h",
						ID:   "4nUgYF7wmkUux1M9SdXM0h",
						Name: "Philharmonic Orchestra",
						Type: "artist",
						URI:  "spotify:artist:4nUgYF7wmkUux1M9SdXM0h",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/66gjGuZX4OGM8CrlfIMftl",
				},
				Href: "https://api.spotify.com/v1/albums/66gjGuZX4OGM8CrlfIMftl",
				ID:   "66gjGuZX4OGM8CrlfIMftl",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b273aed306b9bad470242efb68b8", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e02aed306b9bad470242efb68b8", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d00004851aed306b9bad470242efb68b8", Width: 64},
				},
				Name:                 "bohemian rhapsody - guitar version",
				ReleaseDate:          "2022-11-30",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:66gjGuZX4OGM8CrlfIMftl",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/5Uspdg4EAuwEHmjPtQD8NP",
						},
						Href: "https://api.spotify.com/v1/artists/5Uspdg4EAuwEHmjPtQD8NP",
						ID:   "5Uspdg4EAuwEHmjPtQD8NP",
						Name: "guitar girl",
						Type: "artist",
						URI:  "spotify:artist:5Uspdg4EAuwEHmjPtQD8NP",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/58seMZR64MQxmQHH20LRUQ",
				},
				Href: "https://api.spotify.com/v1/albums/58seMZR64MQxmQHH20LRUQ",
				ID:   "58seMZR64MQxmQHH20LRUQ",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b2738891c122e4f79f262e21ebd5", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e028891c122e4f79f262e21ebd5", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d000048518891c122e4f79f262e21ebd5", Width: 64},
				},
				Name:                 "Bohemian Rhapsody",
				ReleaseDate:          "2020-01-24",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:58seMZR64MQxmQHH20LRUQ",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/0iQDOaYEA5i9RAF0Z73iXb",
						},
						Href: "https://api.spotify.com/v1/artists/0iQDOaYEA5i9RAF0Z73iXb",
						ID:   "0iQDOaYEA5i9RAF0Z73iXb",
						Name: "Angelina Jordan",
						Type: "artist",
						URI:  "spotify:artist:0iQDOaYEA5i9RAF0Z73iXb",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/6mvwcq1i7w481wYCtxJyso",
				},
				Href: "https://api.spotify.com/v1/albums/6mvwcq1i7w481wYCtxJyso",
				ID:   "6mvwcq1i7w481wYCtxJyso",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b273dbf9bbde43127734e1ce9c29", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e02dbf9bbde43127734e1ce9c29", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d00004851dbf9bbde43127734e1ce9c29", Width: 64},
				},
				Name:                 "Bohemian Rhapsody (Acoustic)",
				ReleaseDate:          "2018-10-22",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:6mvwcq1i7w481wYCtxJyso",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/4LaimQU44rsz2kMWQmY6Bi",
						},
						Href: "https://api.spotify.com/v1/artists/4LaimQU44rsz2kMWQmY6Bi",
						ID:   "4LaimQU44rsz2kMWQmY6Bi",
						Name: "John Adams",
						Type: "artist",
						URI:  "spotify:artist:4LaimQU44rsz2kMWQmY6Bi",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/73Td5oIMISZDGAvzGvEpbC",
				},
				Href: "https://api.spotify.com/v1/albums/73Td5oIMISZDGAvzGvEpbC",
				ID:   "73Td5oIMISZDGAvzGvEpbC",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b27340cead1efe274e2658877624", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e0240cead1efe274e2658877624", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d0000485140cead1efe274e2658877624", Width: 64},
				},
				Name:                 "Bohemian Rhapsody",
				ReleaseDate:          "2024-11-12",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:73Td5oIMISZDGAvzGvEpbC",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/32Jb1X3wSmmoHj2epZReZA",
						},
						Href: "https://api.spotify.com/v1/artists/32Jb1X3wSmmoHj2epZReZA",
						ID:   "32Jb1X3wSmmoHj2epZReZA",
						Name: "Steve Vai",
						Type: "artist",
						URI:  "spotify:artist:32Jb1X3wSmmoHj2epZReZA",
					},
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/4IVXzZW6i9KIcpeP29Y1Df",
						},
						Href: "https://api.spotify.com/v1/artists/4IVXzZW6i9KIcpeP29Y1Df",
						ID:   "4IVXzZW6i9KIcpeP29Y1Df",
						Name: "Generation Axe",
						Type: "artist",
						URI:  "spotify:artist:4IVXzZW6i9KIcpeP29Y1Df",
					},
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/2NcbLU1bW55eahD0UgD7U3",
						},
						Href: "https://api.spotify.com/v1/artists/2NcbLU1bW55eahD0UgD7U3",
						ID:   "2NcbLU1bW55eahD0UgD7U3",
						Name: "Brian May",
						Type: "artist",
						URI:  "spotify:artist:2NcbLU1bW55eahD0UgD7U3",
					},
				},
			},
			{
				AlbumType:   "album",
				TotalTracks: 11,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/3LRdN6TlxSk6DJqpSWwdFQ",
				},
				Href: "https://api.spotify.com/v1/albums/3LRdN6TlxSk6DJqpSWwdFQ",
				ID:   "3LRdN6TlxSk6DJqpSWwdFQ",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/64825eca60656397ee6133105402d94df84af353", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/8982055798f8c8dfea1195c539cdc8204925a1e1", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/a8aa505bb39f2369af2eee428281beb3bdf82949", Width: 64},
				},
				Name:                 "Bintang Lima",
				ReleaseDate:          "2000-05-15",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:3LRdN6TlxSk6DJqpSWwdFQ",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/35dWPb7Tmq0WESp6KAUJ8w",
						},
						Href: "https://api.spotify.com/v1/artists/35dWPb7Tmq0WESp6KAUJ8w",
						ID:   "35dWPb7Tmq0WESp6KAUJ8w",
						Name: "Dewa",
						Type: "artist",
						URI:  "spotify:artist:35dWPb7Tmq0WESp6KAUJ8w",
					},
				},
			},
			{
				AlbumType:   "single",
				TotalTracks: 1,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/5S5vnFRSXVua5SbdoY9S8y",
				},
				Href: "https://api.spotify.com/v1/albums/5S5vnFRSXVua5SbdoY9S8y",
				ID:   "5S5vnFRSXVua5SbdoY9S8y",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b2730104b1f71c7bd8e6e5ea1fd2", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e020104b1f71c7bd8e6e5ea1fd2", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d000048510104b1f71c7bd8e6e5ea1fd2", Width: 64},
				},
				Name:                 "Somebody's Pleasure",
				ReleaseDate:          "2023-01-27",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:5S5vnFRSXVua5SbdoY9S8y",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/6ygKuZFz2sRggPZRaLHVHD",
						},
						Href: "https://api.spotify.com/v1/artists/6ygKuZFz2sRggPZRaLHVHD",
						ID:   "6ygKuZFz2sRggPZRaLHVHD",
						Name: "Aziz Hedra",
						Type: "artist",
						URI:  "spotify:artist:6ygKuZFz2sRggPZRaLHVHD",
					},
				},
			},
			{
				AlbumType:   "album",
				TotalTracks: 14,
				ExternalURLs: ExternalURLs{
					Spotify: "https://open.spotify.com/album/2hwNRtZyqmEcCzV47OUAlN",
				},
				Href: "https://api.spotify.com/v1/albums/2hwNRtZyqmEcCzV47OUAlN",
				ID:   "2hwNRtZyqmEcCzV47OUAlN",
				Images: []Image{
					{Height: 640, URL: "https://i.scdn.co/image/ab67616d0000b27369222d0908ecf5d22cc9d960", Width: 640},
					{Height: 300, URL: "https://i.scdn.co/image/ab67616d00001e0269222d0908ecf5d22cc9d960", Width: 300},
					{Height: 64, URL: "https://i.scdn.co/image/ab67616d0000485169222d0908ecf5d22cc9d960", Width: 64},
				},
				Name:                 "The 2000's Greatest",
				ReleaseDate:          "2007-03-12",
				ReleaseDatePrecision: "day",
				Type:                 "album",
				URI:                  "spotify:album:2hwNRtZyqmEcCzV47OUAlN",
				Artists: []Artist{
					{
						ExternalURLs: ExternalURLs{
							Spotify: "https://open.spotify.com/artist/48bKH1ugFBhERC1rdojP9d",
						},
						Href: "https://api.spotify.com/v1/artists/48bKH1ugFBhERC1rdojP9d",
						ID:   "48bKH1ugFBhERC1rdojP9d",
						Name: "Dewa 19",
						Type: "artist",
						URI:  "spotify:artist:48bKH1ugFBhERC1rdojP9d",
					},
				},
			},
		},
	},
}
