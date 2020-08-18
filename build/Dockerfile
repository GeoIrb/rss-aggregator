FROM golang:alpine AS build

WORKDIR /go/src/rss-aggregator
COPY . ./

RUN go build \
    -o /out/service \
    ./cmd/rss-aggregator/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /out/service /app/service
CMD ["/app/service"]