PACKAGES := $(shell go list \
	./cart/internal/service/... \
	./cart/internal/repository/... \
	./cart/test/... \
	| grep -v mock)

run-protoc:
	cd ./cart && make .protoc-generate && cd .. && \
	cd ./loms && make .protoc-generate

run-all:
#	docker-compose up -d
#	docker-compose up --force-recreate --build -d
	docker-compose build --no-cache && docker-compose up --force-recreate -d

run-cover:
	go test -cover $(PACKAGES) | grep -v cart/internal/repository



# for development

run-go-serv:
	go run ./cart/cmd/server/server.go && \
	go run ./loms/cmd/server/server.go

run-doc-base:
	docker-compose up postgres -d && \
	docker-compose up pgadmin -d && \
	docker-compose up prometheus -d && \
	docker-compose up grafana -d

run-doc-dev:
	docker-compose build --no-cache && docker-compose up cart --force-recreate -d && \
	docker-compose up jaeger -d && \
	docker-compose build --no-cache && docker-compose up loms --force-recreate -d 

#	docker-compose build --no-cache && docker-compose up loms --force-recreate -d	

stop-doc-dev:
	docker-compose stop cart && \
	docker-compose stop loms && \
	docker-compose stop jaeger && \
	docker-compose rm -v cart && \
	docker-compose rm -v loms && \
	docker-compose rm -v jaeger
#	 docker volume rm 

stop-doc-all: 
	docker-compose down -v
