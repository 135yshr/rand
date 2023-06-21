all: lint test

install:
	go install github.com/fzipp/gocyclo/cmd/gocyclo
	go install github.com/gordonklaus/ineffassign
	go install github.com/client9/misspell/cmd/misspell

lint: fmt vet staticcheck misspell

fmt:
	gofmt -s -l -w .

vet:
	go vet ./...

staticcheck:
	@[ -x "$(shell which staticcheck)" ] || go install honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...

test:
	go test -cover -v ./...

misspell:
	@[ -x "$(shell which misspell)" ] || go install github.com/client9/misspell/cmd/misspell
	misspell -error ./**/*
