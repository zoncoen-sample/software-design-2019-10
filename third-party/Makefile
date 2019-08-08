.PHONY: build
build: greet en

greet:
	go build -o greet greet.go

en:
	go build -o en ./plugin/en

.PHONY: clean
clean:
	rm -f greet
	rm -f en
