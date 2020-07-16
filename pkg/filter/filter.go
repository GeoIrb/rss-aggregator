package filter

import (
	"time"

	"github.com/GeoIrb/tochka-test/pkg/models"
)

// Filter ...
type Filter interface {
	News(src []models.News, format string, interval time.Duration) (dst []models.News)
}

type filter struct {
}

func (f *filter) News(src []models.News, format string, interval time.Duration) (dst []models.News) {
	dst = make([]models.News, 0, len(src))
	for _, news := range src {
		border := time.Now().Add(-interval)
		pubDate, _ := time.Parse(format, news.PubDate)
		if pubDate.Sub(border) > 0 {
			dst = append(dst, news)
		}
	}
	return
}

// NewFilter ...
func NewFilter() Filter {
	return &filter{}
}
