package rss

import (
	"encoding/xml"

	"github.com/geoirb/rss-aggregator/pkg/models"
)

// RSS ...
type RSS struct {
}

// Parse function for parsing rss-lenta
func (p *RSS) Parse(data []byte) (rss models.Rss) {
	xml.Unmarshal(data, &rss)
	return
}

// NewRSS construct
func NewRSS() *RSS {
	return &RSS{}
}
