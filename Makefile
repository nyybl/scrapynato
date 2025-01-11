build:
	go mod tidy
	go build -v -buildvcs=false -o=build/scrapynato

run:
	./build/scrapynato

reset:
	rm -r build