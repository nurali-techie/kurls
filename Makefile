build: clean
	go build

clean:
	rm -f kurls

clear:
	go mod tidy
	go clean -testcache

test:
	go test ./... -v
