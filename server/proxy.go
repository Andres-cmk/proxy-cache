package server

import (
	"chaching-proxy/cache"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ProxyObject struct {
	Origin string
	Cache  *cache.Cache
	Mutex  sync.RWMutex
}

func NewProxy(origin string) *ProxyObject {
	return &ProxyObject{
		Origin: origin,
		Cache:  cache.NewCache(),
	}
}

func (p *ProxyObject) ClearCache() {
	p.Cache.ClearAll()
	log.Println("Cache cleared")
}

// ServeHTTP maneja las solicitudes HTTP entrantes

func (p *ProxyObject) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("X-Cache")
	header = strings.ToLower(header)
	bodyRequest := r.URL.Query().Get("url")
	if bodyRequest == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	cacheKey := p.Origin + "/" + bodyRequest
	fmt.Println(cacheKey)
	duration := time.Now().Add(10 * time.Minute) // 10 minutos

	switch header {
	case "hit": // from cache
		p.Mutex.Lock()
		if cacheObject, found := p.Cache.Get(cacheKey); found {
			log.Printf("Cache HIT for key: %s", cacheKey)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(cacheObject.ResponseBody)
			if err != nil {
				panic(err.Error())
				return
			}
			if time.Now().After(cacheObject.ExpireAt) {
				p.Cache.ClearRegister(&cacheKey)
				log.Print("Cache cleared for just Expire")
				return
			}
		}
		p.Mutex.Unlock()
		break
	case "miss": // origin server
		log.Printf("Cache MISS for key: %s", cacheKey)
		p.Mutex.Lock()
		response, _ := http.Get(cacheKey)
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)

		cacheObject1 := &cache.CacheObject{
			StatusCode:   response.StatusCode,
			Headers:      header,
			ResponseBody: string(body),
			Created:      time.Now(),
			ExpireAt:     duration,
		}
		p.Cache.Store(&cacheKey, cacheObject1)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(cacheObject1.ResponseBody)
		if err != nil {
			panic(err.Error())
			return
		}
		p.Mutex.Unlock()
		break
	default:
		log.Printf("Unknown header: %s", header)
	}

}
