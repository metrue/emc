clean:
	rm -rf main
build:
	go build -o fibonacci
run: build
	./fibonacci
deploy:
	echo "deploy"
docker:
	docker build -t metrue/fibonacci .
	docker run -p 8080:8080 metrue/fibonacci
