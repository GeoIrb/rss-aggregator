package models

// News ...
type News struct {
	Title string `json:"title" database:"title"`
	Date  string `json:"date" database:"date"`
}
