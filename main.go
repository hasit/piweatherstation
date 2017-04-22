package main

import (
	"strconv"
	"strings"
	"time"

	epdfuse "github.com/wmarbut/go-epdfuse"
)

func main() {
	fuse := epdfuse.NewEpdFuse()

	now := time.Now()

	var mmhh []string
	mmhh = append(mmhh, strconv.Itoa(now.Hour()))
	mmhh = append(mmhh, ":")
	mmhh = append(mmhh, strconv.Itoa(now.Minute()))

	err := fuse.WriteText(strings.Join(mmhh, " "))
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	key := getAPIKey()
// 	lat := "47.7579160"
// 	long := "-122.2396070"

// 	// sample request: https://api.darksky.net/forecast/22f90e2e58fc818dd33419e15933051b/47.7579160,-122.2396070?exclude=minutely,hourly,daily,alerts,flags&units=si
// 	url := "https://api.darksky.net/forecast/" + key + "/" + lat + "," + long + "?exclude=minutely,hourly,daily,alerts,flags&units=si"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	w := new(Weather)
// 	json.NewDecoder(resp.Body).Decode(w)

// 	fmt.Println(w)
// }
