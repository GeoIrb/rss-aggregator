package models

// News ...
type News struct {
	Title   string `db:"title" json:"title" xml:"title"`
	PubDate string `db:"pubDate" json:"pubDate" xml:"pubDate"`
}

// Rss ...
type Rss struct {
	News []News `xml:"channel>item"`
}
