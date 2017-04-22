package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	epdfuse "github.com/wmarbut/go-epdfuse"
)

func main() {
	var t time.Time
	getWeather(t)
	doEvery(1*time.Hour, getWeather)
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func getWeather(t time.Time) {
	key := getAPIKey()
	lat := "47.7579160"
	long := "-122.2396070"

	// sample request: https://api.darksky.net/forecast/22f90e2e58fc818dd33419e15933051b/47.7579160,-122.2396070?exclude=minutely,hourly,daily,alerts,flags&units=si
	url := "https://api.darksky.net/forecast/" + key + "/" + lat + "," + long + "?exclude=minutely,hourly,daily,alerts,flags&units=si"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	w := new(Weather)
	json.NewDecoder(resp.Body).Decode(w)

	var weatherStrings []string

	weatherStrings = append(weatherStrings, w.Currently.Summary)
	tempString := fmt.Sprintf("%.2f", w.Currently.Temperature)
	weatherStrings = append(weatherStrings, tempString)

	fuse := epdfuse.NewEpdFuse()
	err = fuse.Clear()
	if err != nil {
		panic(err)
	}
	err = fuse.WriteText(strings.Join(weatherStrings, "\n"))
	if err != nil {
		panic(err)
	}
}
