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
	AyaName string `json:"aya_name"`
	Text    string `json:"text"`
}
