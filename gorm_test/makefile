NAME = gorm_test

.PHONY: test
build:
	docker build -t ${NAME} -f ./docker/Dockerfile .
up:
	docker-compose -f ./docker/docker-compose.yml up -d
down:
	docker-compose -f ./docker/docker-compose.yml down