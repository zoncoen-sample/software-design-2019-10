.PHONY: build
build: greet en.so jp.so

greet:
	go build -o greet greet.go

en.so:
	go build -buildmode=plugin -o en.so ./plugin/en

jp.so:
	go build -buildmode=plugin -o jp.so ./plugin/jp

.PHONY: clean
clean:
	rm -f greet
	find . -name '*.so' | xargs rm -f
