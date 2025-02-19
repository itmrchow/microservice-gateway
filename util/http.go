package util

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// Try X-Forwarded-For header (used by proxies)
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		// Return the first IP in the X-Forwarded-For chain
		return strings.TrimSpace(ips[0])
	}

	// Try X-Real-IP header
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // If there's an error, return the whole address
	}

	return ip
}
