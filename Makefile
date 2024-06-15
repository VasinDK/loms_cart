PACKAGES := $(shell go list \
	./cart/internal/service/... \
	./cart/internal/repository/... \
	./cart/test/... \
	| grep -v mock)

build-all:
#	cd cart && GOOS=linux GOARCH=amd64 make build


run-all: build-all
	docker-compose up --force-recreate --build -d

run-cover:
	go test -cover $(PACKAGES) | grep -v cart/internal/repository

run-loms:
	cd ./loms && make .protoc-generate