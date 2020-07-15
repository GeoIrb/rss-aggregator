// Package httpserver ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpserver

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/GeoIrb/tochka-test/pkg/models"
	"github.com/valyala/fasthttp"
)

var (
	emptyBytes = []byte("")
)

type startTrackingRequest struct {
	Format string `json:"format"`
	URL    string `json:"url"`
}

// StartTrackingTransport transport interface
type StartTrackingTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (url string, format string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type startTrackingTransport struct {
	decodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *startTrackingTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (url string, format string, err error) {
	var request startTrackingRequest
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		err = t.decodeJSONErrorCreator(http.StatusBadRequest, err)
		return
	}
	format = request.Format
	url = request.URL
	return
}

// EncodeResponse method for encoding response on server side
func (t *startTrackingTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.SetStatusCode(http.StatusCreated)
	return
}

// NewStartTrackingTransport the transport creator for http requests
func NewStartTrackingTransport(decodeJSONErrorCreator errorCreator) StartTrackingTransport {
	return &startTrackingTransport{
		decodeJSONErrorCreator: decodeJSONErrorCreator,
	}
}

type stopTrackingRequest struct {
	URL string `json:"url"`
}

// StopTrackingTransport transport interface
type StopTrackingTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (url string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error)
}

type stopTrackingTransport struct {
	decodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *stopTrackingTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (url string, err error) {
	var request stopTrackingRequest
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		err = t.decodeJSONErrorCreator(http.StatusBadRequest, err)
		return
	}
	url = request.URL
	return
}

// EncodeResponse method for encoding response on server side
func (t *stopTrackingTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response) (err error) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewStopTrackingTransport the transport creator for http requests
func NewStopTrackingTransport(decodeJSONErrorCreator errorCreator) StopTrackingTransport {
	return &stopTrackingTransport{
		decodeJSONErrorCreator: decodeJSONErrorCreator,
	}
}

//easyjson:json
type getNewsResponse struct {
	News []models.News `json:"news"`
}

// GetNewsTransport transport interface
type GetNewsTransport interface {
	DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (title *string, err error)
	EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, news []models.News) (err error)
}

type getNewsTransport struct {
	encodeJSONErrorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getNewsTransport) DecodeRequest(ctx *fasthttp.RequestCtx, r *fasthttp.Request) (title *string, err error) {
	title = ptr(ctx.QueryArgs().Peek("title"))
	return
}

// EncodeResponse method for encoding response on server side
func (t *getNewsTransport) EncodeResponse(ctx *fasthttp.RequestCtx, r *fasthttp.Response, news []models.News) (err error) {
	r.Header.Set("Content-Type", "application/json")
	var theResponse getNewsResponse
	theResponse.News = news
	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = t.encodeJSONErrorCreator(http.StatusInternalServerError, err)
		return
	}
	r.SetBody(body)
	r.Header.SetStatusCode(http.StatusOK)
	return
}

// NewGetNewsTransport the transport creator for http requests
func NewGetNewsTransport(encodeJSONErrorCreator errorCreator) GetNewsTransport {
	return &getNewsTransport{
		encodeJSONErrorCreator: encodeJSONErrorCreator,
	}
}

func ptr(in []byte) *string {
	i := string(in)
	return &i
}

func atoiptr(in []byte) (out *int, err error) {
	var (
		o int
		i = string(in)
	)
	if i != "" {
		if o, err = strconv.Atoi(i); err == nil {
			out = &o
		}
	}
	return
}

func atoi(in []byte) (out int, err error) {
	if bytes.Equal(in, emptyBytes) {
		return
	}
	return strconv.Atoi(string(in))
}
