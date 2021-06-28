GO_VERSION := 1.16.5

setup: install-go init-go

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz" --no-check-certificate
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

build:
	go build -o api main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html