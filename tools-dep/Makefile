tools:
	@export GOBIN=$(PWD)/bin
	go install github.com/vektra/mockery/v2

generate: tools
	go generate ./...