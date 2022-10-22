package main

type Token struct {
	ConsumerKey    string
	ConsumerSecret string
	ApiKey         string
	ApiSecret      string
}

type ResponseTweet struct {
	Status string `json:"status"`
	Acak   struct {
		Id RandomAyat `json:"id"`
	} `json:"acak"`
}
type RandomAyat struct {
	Id    string `json:"id"`
	Surat string `json:"surat"`
	Ayat  string `json:"ayat"`
	Teks  string `json:"teks"`
}
