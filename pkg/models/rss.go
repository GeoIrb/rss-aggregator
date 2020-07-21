package models

// News ...
type News struct {
	Title   string `db:"title" xml:"title" json:"title"`
	PubDate string `db:"pubDate" xml:"pubDate" json:"pubDate"`
}

// Rss struct rss-line
type Rss struct {
	News []News `xml:"channel>item" json:"-"`
}
