BINARY=terraform-ui

all: build testing

build:
	@echo "~RUNNING BUILD~"
	go build -o $(BINARY) cmd/terraform-ui/main.go

testing:
	@echo "~RUNNING TEST~"
	go test --cover ./...