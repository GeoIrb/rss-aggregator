package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"

	"github.com/GeoIrb/rss-aggregator/pkg/converter"
	"github.com/GeoIrb/rss-aggregator/pkg/filter"
	"github.com/GeoIrb/rss-aggregator/pkg/postgres"
	"github.com/GeoIrb/rss-aggregator/pkg/rss"
	"github.com/GeoIrb/rss-aggregator/pkg/service"
	"github.com/GeoIrb/rss-aggregator/pkg/service/httperrors"
	"github.com/GeoIrb/rss-aggregator/pkg/service/httpserver"
	"github.com/GeoIrb/rss-aggregator/pkg/site"
	"github.com/GeoIrb/rss-aggregator/pkg/storage"
)

type configuration struct {
	Port string `envconfig:"PORT" default:"8080"`

	DBHost          string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort          int    `envconfig:"DB_PORT" default:"5432"`
	DBUser          string `envconfig:"DB_USER" default:"tochka"`
	DBPassword      string `envconfig:"DB_PASSWORD" default:"tochka"`
	DBName          string `envconfig:"DB_NAME" default:"tochka"`
	DBConnectDriver string `envconfig:"DB_CONNECT_DRIVER" default:"postgres"`
	DBConnectLayout string `envconfig:"DB_CONNECT_LAYOUT" default:"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"`

	StorageInsertNews   string `envconfig:"STORAGE_INSERT_NEWS" default:"INSERT INTO public.news (\"title\", \"pubDate\") VALUES ($1, $2)"`
	StorageSelectNews   string `envconfig:"STORAGE_SELECT_NEWS" default:"SELECT \"title\", \"pubDate\" FROM public.news WHERE title like '%' || $1 || '%'"`
	StorageAllTitleNews string `envconfig:"STORAGE_ALL_TITLE_NEWS" default:""`

	IntervalTracking time.Duration `envconfig:"INTERVAL_TRACKING" default:"10m"`
	SiteTimeout      time.Duration `envconfig:"SITE_TIMEOUT" default:"2s"`
}

func main() {
	var cfg configuration

	if err := envconfig.Process("", &cfg); err != nil {
		os.Exit(1)
	}

	st := site.NewSite(cfg.SiteTimeout)
	r := rss.NewRSS()
	flr := filter.NewFilter()
	cnv := converter.NewConverter()

	db := postgres.NewPostgres(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,

		cfg.DBConnectDriver,
		cfg.DBConnectLayout,
	)

	if err := db.Connect(); err != nil {
		fmt.Printf("failed to connect database %s", err)
		os.Exit(1)
	}

	srg := storage.NewStorage(
		db,
		cfg.StorageInsertNews,
		cfg.StorageSelectNews,
		cfg.StorageAllTitleNews,
	)

	svc := service.NewService(
		st,
		r,
		flr,
		cnv,
		srg,
		cfg.IntervalTracking,
	)

	svc.Start()
	defer svc.Shoutdown()

	errorProccessor := httperrors.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	router := fasthttprouter.New()
	httpserver.New(router, svc, httperrors.NewError, httperrors.NewError, errorProccessor)

	server := &fasthttp.Server{
		Handler: router.Handler,
	}

	go func() {
		fmt.Printf("start server port: %s", cfg.Port)
		if err := server.ListenAndServe(":" + cfg.Port); err != nil {
			fmt.Printf("server run failure error %s", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	
	sig := <-c
	fmt.Printf("received signal, exiting signal %v", sig)

	if err := server.Shutdown(); err != nil {
		fmt.Printf("server shutdown failure %v", err)
	}

	fmt.Printf("goodbye")
}
