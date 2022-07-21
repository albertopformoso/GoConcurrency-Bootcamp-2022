build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o main ./main.go

redis:
	docker-compose -f ./docker/docker-compose.yml up -d --build redis

rm-redis:
	docker-compose -f ./docker/docker-compose.yml down

clean: rm-redis
	docker rmi -f redis