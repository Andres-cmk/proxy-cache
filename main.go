package main

import (
	"chaching-proxy/server"
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

	proxyServer := server.NewProxy(*urlTarget)

	http.HandleFunc("/", proxyServer.HandleHTTP)

	log.Println("Starting proxy server: ", proxyServer)

	if *clearCache {
		proxyServer.ClearCache()
	}

	log.Print("Listening on port " + *port + "...")
	_ = http.ListenAndServe(":"+*port, nil)

}
