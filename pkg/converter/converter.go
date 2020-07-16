package converter

import (
	"github.com/GeoIrb/tochka-test/pkg/models"
)

// Converter ...
type Converter interface {
	News(src []models.News) (dst [][]string)
}

type converter struct {
}

func (c *converter) News(src []models.News) (dst [][]string) {
	dst = make([][]string, 0, len(src))
	for _, news := range src {
		data := make([]string, 2)
		data[0] = news.Title
		data[1] = news.PubDate
		dst = append(dst, data)
	}
	return
}

// NewConverter ...
func NewConverter() Converter {
	return &converter{}
}
