package main

import (
	"flag"
	"fmt"
	"net/http"
)

var endpoint, version, key string

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "Microsoft client endpoint")
	flag.StringVar(&version, "version", "v2.0", "Microsoft client version")
	flag.StringVar(&key, "key", "", "Microsoft client key")
	flag.Parse()
}

func main() {
	api := APIClient{
		endpoint,
		version,
		key,
		&http.Client{},
	}
	sentimentResponse := api.sentiment("en", 1, "Hello, World!")
	fmt.Printf("Score: %f\n", sentimentResponse.Documents[0].Score)
}
