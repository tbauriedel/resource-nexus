BINARY=resource-nexus-core

all: build testing cleanup

build:
	@echo "~RUNNING BUILD~"
	go build -o $(BINARY) cmd/resource-nexus-core/main.go

testing:
	@echo "~RUNNING TEST~"
	go test --cover ./...

cleanup:
	@echo "~RUNNING CLEANUP~"
	rm -f ./resource-nexus-core