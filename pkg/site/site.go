package site

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Site handler
type Site struct {
	timeout time.Duration
}

// GetDatа function for get data from url
func (s *Site) GetDatа(url string) (data []byte, err error) {
	_, data, err = fasthttp.GetTimeout(nil, url, s.timeout)
	return
}

// NewSite constructor site handler
func NewSite(timeout time.Duration) *Site {
	return &Site{
		timeout: timeout,
	}
}
