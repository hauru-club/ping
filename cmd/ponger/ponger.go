package main

import (
	"log"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"

	"github.com/hauru-club/ping/assets"
	"github.com/hauru-club/ping/pkg/handlers"
)

func main() {
	s := sse.NewServer(nil)
	defer s.Shutdown()

	http.Handle("/", handlers.Bytes(assets.Index))
	http.Handle("/events/", s)
	http.Handle("/publish", handlers.Publish(s))
	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.FS(assets.StaticFiles))),
	)
	log.Println("Listening at :5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
