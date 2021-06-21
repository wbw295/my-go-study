package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()

	http.Handle("/ws", Middleware(
		http.HandlerFunc(wsHandler),
		authMiddleware,
	))
	log.Printf("listening on %v", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(rw http.ResponseWriter, req *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Printf("upgrade err: %v", err)
		return
	}
	defer wsConn.Close()

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Printf("read err: %v", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	TestApiKey := "test_api_key"
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var apiKey string
		if apiKey = req.Header.Get("X-Api-Key"); apiKey != TestApiKey {
			log.Printf("bad auth api key: %s", apiKey)
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(rw, req)
	})
}