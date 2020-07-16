// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"

	"github.com/GeoIrb/rss-aggregator/pkg/models"
	"github.com/valyala/fasthttp"
)

type service interface {
	StartTracking(ctx context.Context, url string, format string) (err error)
	StopTracking(ctx context.Context, url string) (err error)
	GetNews(ctx context.Context, title *string) (news []models.News, err error)
}

type startTrackingSwaggerInfo struct {
	transport      StartTrackingTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *startTrackingSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		url    string
		format string
		err    error
	)
	url, format, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.StartTracking(ctx, url, format)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewStartTrackingSwaggerInfo the server creator
func NewStartTrackingSwaggerInfo(transport StartTrackingTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := startTrackingSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type stopTrackingSwaggerInfo struct {
	transport      StopTrackingTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *stopTrackingSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		url string
		err error
	)
	url, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	err = s.service.StopTracking(ctx, url)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewStopTrackingSwaggerInfo the server creator
func NewStopTrackingSwaggerInfo(transport StopTrackingTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := stopTrackingSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getNewsSwaggerInfo struct {
	transport      GetNewsTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getNewsSwaggerInfo) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		title *string
		news  []models.News
		err   error
	)
	title, err = s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	news, err = s.service.GetNews(ctx, title)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err = s.transport.EncodeResponse(ctx, &ctx.Response, news); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetNewsSwaggerInfo the server creator
func NewGetNewsSwaggerInfo(transport GetNewsTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getNewsSwaggerInfo{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
