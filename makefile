.PHONY: build

build: 
	go build -o ./simple_website ./cmd/.

.PHONY: run

run: build
	./simple_website


