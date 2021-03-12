package signaling

import (
	"os"
)

func URI() string {
	defaultURL := "http://localhost:8080"
	if url := os.Getenv("SIGNALING_URL"); url != "" {
		return url
	}
	return defaultURL
}

// ConnectInfo SDP by offer or answer
type ConnectInfo struct {
	Source string `json:"source"`
	SDP    string `json:"sdp"`
}
