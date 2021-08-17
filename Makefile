GO_VERSION := 1.16.5
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse HEAD)
DATE := $(shell date +%Y-%m-%d.%H:%M:%S)
LDFLAGS := -w -X github.com/holmes89/hello-api/handlers.hash=$(HASH) -X github.com/holmes89/hello-api/handlers.tag=$(TAG) -X github.com/holmes89/hello-api/handlers.date=$(DATE)

setup: install-go init-go isntall-lint copy-hooks

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz" --no-check-certificate
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

ldflags:
	@echo $(LDFLAGS)

build: 
	go build -ldflags "$(LDFLAGS)" -o api cmd/paas/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out \
	| grep "total:" | awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1

check-format:
	test -z $$(go fmt ./...)

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.