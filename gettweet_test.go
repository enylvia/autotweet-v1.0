package main

import (
	"encoding/json"
	"github.com/rivo/uniseg"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestGetTweet(t *testing.T) {
	// Test Get Tweet
	var response ResponseTweet
	random := 1
	resp, _ := http.Get("https://api.myquran.com/v1/tafsir/quran/kemenag/id/" + strconv.Itoa(random))
	//limit character
	defer resp.Body.Close()
	// Decode data from API
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
	}
	text := response.Data[0].AyaName + " - " + response.Data[0].Text
	count := uniseg.GraphemeClusterCount(text)
	if count > 275 {
		text = LimitCharacter(response.Data[0].AyaName+" - "+response.Data[0].Text, 270)
	}
	log.Println(count)
	log.Println(text)

}
