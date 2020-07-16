package rss

import (
	"encoding/xml"

	"github.com/GeoIrb/rss-aggregator/pkg/models"
)

// RSS ...
type RSS interface {
	Parse(data []byte) (rss models.Rss)
}

type rss struct {
}

func (p *rss) Parse(data []byte) (rss models.Rss) {
	xml.Unmarshal(data, &rss)
	return
}

// NewRSS ...
func NewRSS() RSS {
	return &rss{}
}
