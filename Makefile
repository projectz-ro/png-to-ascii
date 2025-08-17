.PHONY: build clean

build: 
	mkdir -p build
	go build -o build/png-to-ascii ./cmd/*.go

clean:
	rm -rf build
