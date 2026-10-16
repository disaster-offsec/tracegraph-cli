package config

import "os"

// defaultBackend is used when neither a flag nor an env var is set.
const defaultBackend = "http://localhost:8000"

// BackendURL returns the backend address.
// Priority: flag > env > default.
// An empty flagValue means the flag was not provided.
func BackendURL(flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	if v := os.Getenv("TRACEGRAPH_BACKEND"); v != "" {
		return v
	}
	return defaultBackend
}
