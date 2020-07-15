// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"context"
	"net/http/pprof"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodStartTracking = "POST"
	uriPathStartTracking    = "/tracking"
	httpMethodStopTracking  = "DELETE"
	uriPathStopTracking     = "/tracking"
	httpMethodGetNews       = "GET"
	uriPathGetNews          = "/news"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type errorCreator func(code int, err error) error

// New ...
func New(router *fasthttprouter.Router, svc service, decodeJSONErrorCreator errorCreator, encodeJSONErrorCreator errorCreator, errorProcessor errorProcessor) {

	startTrackingTransport := NewStartTrackingTransport(decodeJSONErrorCreator)
	router.Handle(httpMethodStartTracking, uriPathStartTracking, NewStartTrackingSwaggerInfo(startTrackingTransport, svc, errorProcessor))

	stopTrackingTransport := NewStopTrackingTransport(decodeJSONErrorCreator)
	router.Handle(httpMethodStopTracking, uriPathStopTracking, NewStopTrackingSwaggerInfo(stopTrackingTransport, svc, errorProcessor))

	getNewsTransport := NewGetNewsTransport(encodeJSONErrorCreator)
	router.Handle(httpMethodGetNews, uriPathGetNews, NewGetNewsSwaggerInfo(getNewsTransport, svc, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
