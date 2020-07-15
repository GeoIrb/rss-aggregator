package site

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Site ...
type Site interface {
	GetDate(url string) (data []byte, err error)
}

type site struct {
	timeout time.Duration
}

func (s *site) GetDate(url string) (data []byte, err error) {
	_, data, err = fasthttp.GetTimeout(nil, url, s.timeout)
	return
}

// NewSite ...
func NewSite(timeout time.Duration) Site {
	return &site{
		timeout: timeout,
	}
}
