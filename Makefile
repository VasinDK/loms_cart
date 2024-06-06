build-all:
#	cd cart && GOOS=linux GOARCH=amd64 make build


run-all: build-all
	docker-compose up --force-recreate --build -d

run-cover:
	All_PKG=$(shell go list ./cart/internal/service/...)
	echo "$All_PKG"
#	pwd 
#	cd ./cart