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

run:
	go run ./cart/cmd/server/server.go

run-docker:
	docker-compose up postgres -d && \
	docker-compose build --no-cache && docker-compose up cart --force-recreate -d && \
	docker-compose up pgadmin -d && \
	docker-compose up prometheus -d && \
	docker-compose up jaeger -d && \
	docker-compose build --no-cache && docker-compose up loms --force-recreate -d	
#	docker-compose up loms --force-recreate --build -d

stop-docker: 
	docker-compose down -v
