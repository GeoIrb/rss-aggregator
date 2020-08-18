check-code:
	go get -u golang.org/x/lint/golint
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go fmt ./...
	go vet ./...
	out=$(go fmt ./...) && if [[ -n "$out" ]]; then echo "$out"; exit 1; fi
	golint -set_exit_status $(go list ./...)
	golangci-lint run -E gofmt -E golint -E vet

run:
	docker-compose -f build/docker-compose.yml -d up