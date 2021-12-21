build: install-tools generate-sources

install-tools:
	go install $$(go list -f '{{join .Imports " "}}' tools.go)

generate-sources: 
	go generate ./...

compile-sources:
	go build ./...
