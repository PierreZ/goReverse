package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func main() {

	var destURL *url.URL

	// Def & parsing of flags
	originPtr := flag.Int("p", 8080, "The proxy listening port")
	destPtr := flag.String("d", "", "The destination")
	flag.Parse()

	destURL, err := url.Parse(*destPtr)
	if err != nil {
		log.Fatal("parse url: ", err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(destURL)

	log.Println("Starting goReverse")

	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(*originPtr), 10), reverseProxy))
}
