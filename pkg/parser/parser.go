package parser

import (
	"encoding/xml"
	"time"

	"github.com/GeoIrb/tochka-test/pkg/models"
)

// Parser ...
type Parser interface {
	News(data []byte, format string, interval time.Duration) (news [][]string)
}

type parser struct {
}

func (p *parser) News(data []byte, format string, interval time.Duration) (news [][]string) {
	var lenta models.Rss
	if err := xml.Unmarshal(data, &lenta); err != nil {
		return nil
	}

	for _, item := range lenta.Items {
		border := time.Now().Add(-interval)
		pubDate, _ := time.Parse(format, item.PubDate)
		if pubDate.Sub(border) > 0 {
			itm := make([]string, 2)
			itm[0] = item.Title
			itm[1] = item.PubDate
			news = append(news, itm)
		}
	}
	return
}

// NewParser ...
func NewParser() Parser {
	return &parser{}
}
