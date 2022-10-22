package main

type Token struct {
	ConsumerKey    string
	ConsumerSecret string
	ApiKey         string
	ApiSecret      string
}

type ResponseTweet struct {
	Status bool         `json:"status"`
	Data   []RandomAyat `json:"data"`
}
type RandomAyat struct {
	NoSurah      int    `json:"no_surah"`
	NoAyat       int    `json:"no_ayat"`
	TeksTerjemah string `json:"teks_terjemah"`
}
