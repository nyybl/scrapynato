build:
	go mod tidy
	go build -v -buildvcs=false -o=build/scrapynato

start:
	./build/scrapynato

dev:
	go run .

reset:
	rm -r build