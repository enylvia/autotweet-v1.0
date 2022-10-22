package main

import (
	"encoding/json"
	"fmt"
	"github.com/rivo/uniseg"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestGetTweet(t *testing.T) {
	// Test Get Tweet
	var response ResponseTweet
	surah := RandomNumberGiven(1, 77)
	aya := RandomNumberGiven(1, 11)

	formatUrl := fmt.Sprintf("https://quran.kemenag.go.id/api/v1/ayatweb/%d/%d/%d/%d", surah, aya, 0, 6236)
	fmt.Println(surah)
	fmt.Println(aya)
	resp, err := http.Get(formatUrl)
	//limit character
	defer resp.Body.Close()
	// Decode data from API
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
	text := strconv.Itoa(response.Data[0].NoSurah) + " - " + response.Data[0].TeksTerjemah
	count := uniseg.GraphemeClusterCount(text)
	if count > 275 {
		text = LimitCharacter(text, 270)
	}
	log.Println(count)
	log.Println(text)

}
