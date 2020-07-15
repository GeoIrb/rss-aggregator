package models

// Item ...
type Item struct {
	PubDate string `xml:"pubDate"`
	Title   string `xml:"title"`
}

// Rss ...
type Rss struct {
	Items []Item `xml:"channel>item"`
}
