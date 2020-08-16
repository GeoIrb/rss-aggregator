package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/geoirb/rss-aggregator/pkg/models"
)

type storage interface {
	AddNews(newsField ...string) (err error)
	GetNews(ctx context.Context, title *string) (news []models.News, err error)
}

type source interface {
	GetDatа(url string) (data []byte, err error)
}

type rss interface {
	Parse(data []byte) (news models.Rss)
}

type filter interface {
	News(src []models.News, format string, interval time.Duration) (dst []models.News)
}

type converter interface {
	News(src []models.News) (dst [][]string)
}

// Service ...
// @gtg http-server http-errors
type Service interface {
	Start()
	Shoutdown()
	// @gtg http-server-method POST
	// @gtg http-server-uri-path /tracking
	// @gtg http-server-json-tag url url
	// @gtg http-server-json-tag format format
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusCreated
	// @gtg http-server-response-content-type application/json
	StartTracking(ctx context.Context, url string, format string) (err error)
	// @gtg http-server-method DELETE
	// @gtg http-server-uri-path /tracking
	// @gtg http-server-json-tag url url
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	StopTracking(ctx context.Context, url string) (err error)
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /news
	// @gtg http-server-query title={title}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-json-tag news news
	// @gtg http-server-response-content-type application/json
	GetNews(ctx context.Context, title *string) (news []models.News, err error)
}

type service struct {
	source    source
	rss       rss
	filter    filter
	converter converter
	storage   storage

	cache    sync.Map
	interval time.Duration
	newsChan chan [][]string
}

func (s *service) Start() {
	go func() {
		for news := range s.newsChan {
			for _, n := range news {
				s.storage.AddNews(n...)
			}
		}
	}()
}

func (s *service) Shoutdown() {
	s.cache.Range(func(k, v interface{}) bool {
		v.(context.CancelFunc)()
		return true
	})
	close(s.newsChan)
}

func (s *service) StartTracking(ctx context.Context, url, format string) (err error) {
	if _, isExist := s.cache.Load(url); isExist {
		err = fmt.Errorf("%s is exist", url)
		return
	}

	trackingCtx, trackingCnl := context.WithCancel(context.Background())
	go s.tracking(trackingCtx, url, format)

	s.cache.Store(url, trackingCnl)
	return
}

func (s *service) StopTracking(ctx context.Context, url string) (err error) {
	var (
		cancel  interface{}
		isExist bool
	)
	if cancel, isExist = s.cache.Load(url); !isExist {
		err = fmt.Errorf("%s is not exist", url)
		return
	}

	cancel.(context.CancelFunc)()
	s.cache.Delete(url)
	return
}

func (s *service) GetNews(ctx context.Context, title *string) (news []models.News, err error) {
	news, err = s.storage.GetNews(ctx, title)
	return
}

func (s *service) tracking(ctx context.Context, url, format string) {
	t := time.NewTicker(s.interval)

	s.getNews(url, format)
	for {
		select {
		case <-t.C:
			s.getNews(url, format)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}
}

func (s *service) getNews(url, format string) {
	if data, err := s.source.GetDatа(url); err == nil {
		rss := s.rss.Parse(data)
		news := s.filter.News(rss.News, format, s.interval)
		s.newsChan <- s.converter.News(news)
	}
}

// NewService ...
func NewService(
	source source,
	rss rss,
	filter filter,
	converter converter,
	storage storage,
	interval time.Duration,
) Service {
	return &service{
		source:    source,
		rss:       rss,
		filter:    filter,
		converter: converter,
		storage:   storage,
		interval:  interval,

		newsChan: make(chan [][]string, 1),
	}
}
