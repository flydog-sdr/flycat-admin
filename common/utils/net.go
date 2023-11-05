package utils

import (
	"net/http"
	"time"
)

func HasInternet(timeout time.Duration) bool {
	client := http.Client{
		Timeout: timeout,
	}

	_, err := client.Get("http://cp.cloudflare.com")
	return err == nil
}
