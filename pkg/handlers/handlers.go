// Package handlers implements http handlers for ping http server.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/hauru-club/ping/assets"
	"github.com/hauru-club/ping/pkg/message"
	"github.com/hauru-club/ping/pkg/models"
)

func methodNotAllowed(w http.ResponseWriter, method string) {
	http.Error(w,
		fmt.Sprintf("Method %s is not allowed", method),
		http.StatusMethodNotAllowed,
	)
}

func internalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)

}

// Publish is handler for publishing ICMP packets to
// SSE clients.
func Publish(s *sse.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			methodNotAllowed(w, r.Method)
			return
		}

		p := new(models.Packet)
		if err := json.NewDecoder(r.Body).Decode(p); err != nil {
			log.Printf("Publish error: %s", err)
			internalServerError(w)
			return
		}

		s.SendMessage("/events/test", message.JSON(p, ""))
	}
}

// Bytes outputs given slice of bytes to http client
// with http Status OK.
func Bytes(index []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			methodNotAllowed(w, r.Method)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(assets.Index)
	}
}
