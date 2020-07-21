package storage

import (
	"context"
	"fmt"

	"github.com/GeoIrb/rss-aggregator/pkg/models"
)

type db interface {
	Connect() (err error)
	Ping() (err error)
	Query(query string, args ...interface{}) (err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
}

// Storage ...
type Storage struct {
	db db

	insertNews string
	selectNews string

	allTitle string
}

// AddNews add news in storage
func (s *Storage) AddNews(newsField ...string) (err error) {
	if err = s.db.Ping(); err != nil {
		if err = s.db.Connect(); err != nil {
			err = fmt.Errorf("connect db %s", err)
			return
		}
	}

	params := make([]interface{}, len(newsField))
	for i, v := range newsField {
		params[i] = v
	}

	if err = s.db.Query(s.insertNews, params...); err != nil {
		err = fmt.Errorf("query db %s", err)
	}
	return
}

// GetNews get news form storage
func (s *Storage) GetNews(ctx context.Context, title *string) (news []models.News, err error) {
	if err = s.db.Ping(); err != nil {
		if err = s.db.Connect(); err != nil {
			err = fmt.Errorf("connect db %s", err)
			return
		}
	}

	findTitle := s.allTitle
	if title != nil {
		findTitle = *title
	}

	if err = s.db.Select(&news, s.selectNews, findTitle); err != nil {
		err = fmt.Errorf("query db %s", err)
	}
	return
}

// NewStorage construct 
func NewStorage(
	db db,

	insertNews string,
	selectNews string,

	allTitle string,
) *Storage {
	return &Storage{
		db:         db,
		insertNews: insertNews,
		selectNews: selectNews,
		allTitle:   allTitle,
	}
}
