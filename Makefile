build:
	go build

install:
	go install

docker: docker-build docker-push

docker-build:
	docker build -t ondrejsika/loggen .

docker-push:
	docker push ondrejsika/loggen
