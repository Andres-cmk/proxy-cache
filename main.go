package main

import (
	"chaching-proxy/cache"
	"flag"
	"log"
	"net/http"
)

func main() {
	// inputs por la linea de comandos.
	port := flag.String("port", "", "-port")
	urlTarget := flag.String("origin", "", "-url")
	clearCache := flag.Bool("clear-cache", false, "-clear-cache")
	flag.Parse()
	StartServer(port, clearCache, urlTarget)
}

func StartServer(port *string, clearCache *bool, target *string) {
	newCache := cache.NewCache()
	if *clearCache == true {
		newCache.ClearAll()
	}
	log.Print("Listening on port " + *port + "...")
	_ = http.ListenAndServe(":"+*port, nil)
}
