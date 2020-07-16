package models

// News ...
type News struct {
	Title   string `db:"title" xml:"title" json:"title"`
	PubDate string `db:"pubDate" xml:"pubDate" json:"pubDate"`
}

// Rss ...
type Rss struct {
	News []News `xml:"channel>item" json:"-"`
}
