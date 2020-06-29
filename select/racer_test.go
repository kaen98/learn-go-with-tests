package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.sina.com"
	fastURL := "http://www.baidu.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got , want)
	}
}