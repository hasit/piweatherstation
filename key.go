package main

import "os"

func getAPIKey() string {
	key := os.Getenv("DARKSKY_SECRETKEY")

	if len(key) == 0 {
		panic("DARKSKY_SECRETKEY environment variable is not set!\n")
	}

	return key
}
