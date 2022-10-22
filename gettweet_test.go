package main

import (
	"github.com/rivo/uniseg"
	"testing"
)

func TestGetTweet(t *testing.T) {
	// Test Get Tweet
	data, err := GetTweet()
	if err != nil {
		t.Error("Error Get Tweet")
	}
	count := uniseg.GraphemeClusterCount(data.Acak.Id.Teks)
	if count > 260 {
		GetTweet()
	}
	if count < 260 {
		t.Log("Success Get Tweet")
	}

}
