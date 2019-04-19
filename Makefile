clean:
	rm -rf main
build:
	go build main.go handler.go cache.go
run: build
	./main
deploy:
	echo "deploy"
