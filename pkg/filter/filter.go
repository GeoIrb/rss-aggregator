package filter

import (
	"time"

	"github.com/GeoIrb/rss-aggregator/pkg/models"
)

// Filter ...
type Filter struct {
}

// News filtering news
func (f *Filter) News(src []models.News, format string, interval time.Duration) (dst []models.News) {
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

// NewFilter construct
func NewFilter() *Filter {
	return &Filter{}
}
