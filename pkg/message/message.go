// Package message contains custom SSE messages.
package message

import (
	"bytes"
	"encoding/json"

	"github.com/alexandrevicenzi/go-sse"
)

// JSON returns SSE message with json string representation
// of given data.
func JSON(data interface{}, event string) *sse.Message {
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(data); err != nil {
		return sse.NewMessage("", "", "")
	}
	return sse.NewMessage("", buffer.String(), event)
}
