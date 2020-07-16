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
type Storage interface {
	AddNews(newsField ...string) (err error)
	GetNews(ctx context.Context, title *string) (news []models.News, err error)
}

type storage struct {
	db db

	insertNews string
	selectNews string

	allTitle string
}

func (s *storage) AddNews(newsField ...string) (err error) {
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

func (s *storage) GetNews(ctx context.Context, title *string) (news []models.News, err error) {
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

// NewStorage ...
func NewStorage(
	db db,

	insertNews string,
	selectNews string,

	allTitle string,
) Storage {
	return &storage{
		db:         db,
		insertNews: insertNews,
		selectNews: selectNews,
		allTitle:   allTitle,
	}
}
