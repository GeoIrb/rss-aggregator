package converter

import (
	"github.com/geoirb/rss-aggregator/pkg/models"
)

// Converter ...
type Converter struct {
}

// News convert news to slice of []string
func (c *Converter) News(src []models.News) (dst [][]string) {
	dst = make([][]string, 0, len(src))
	for _, news := range src {
		data := make([]string, 2)
		data[0] = news.Title
		data[1] = news.PubDate
		dst = append(dst, data)
	}
	return
}

// NewConverter construct
func NewConverter() *Converter {
	return &Converter{}
}
