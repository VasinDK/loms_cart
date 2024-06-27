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

stop-docker: 
	docker-compose down

run-cover:
	go test -cover $(PACKAGES) | grep -v cart/internal/repository

# run-migrations:
#	goose -dir ./loms/migrations postgres "postgresql://admin_loms:password@localhost:5432/loms?sslmode=disable" up

run-cart:
	go run ./cart/cmd/server/server.go