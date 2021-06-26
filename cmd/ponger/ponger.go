package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alexandrevicenzi/go-sse"

	"github.com/hauru-club/ping/assets"
	"github.com/hauru-club/ping/pkg/handlers"
)

const (
	secretKeyEnv          = "PONGER_SECRET"
	defaultSecretKeyValue = "supersecretkey"

	portEnv          = "PONGER_PORT"
	defaultPortValue = "5000"

	addressEnv          = "PONGER_ADDRESS"
	defaultAddressValue = "0.0.0.0"
)

func readSetting(env, fallback string) string {
	res := os.Getenv(env)
	if res == "" {
		return fallback
	}
	return res
}

func main() {
	secretKey := readSetting(secretKeyEnv, defaultSecretKeyValue)
	port := readSetting(portEnv, defaultPortValue)
	address := readSetting(addressEnv, defaultAddressValue)

	s := sse.NewServer(nil)
	defer s.Shutdown()

	http.Handle("/", handlers.Bytes(assets.Index))
	http.Handle("/events/", s)

	http.Handle("/publish", handlers.AuthKeyMiddleware(secretKey)(handlers.Publish(s)))
	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.FS(assets.StaticFiles))),
	)
	log.Println("Listening at", address+":"+port)
	http.ListenAndServe(address+":"+port, nil)
}
