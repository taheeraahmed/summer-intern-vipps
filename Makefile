swagger: 
	swagger generate server -f ./manifests/k8s/base/openApiSwagger/summerstudents-backend-apim.yaml -t generated -A summerstudents-backend --main-package ../../cmd && go mod tidy

run:
	@git ls-files "**/*.go" | entr -r env -S "`cat $(wildcard .env)`" go run cmd/main.go
tidy:
	go mod tidy

docker:
	docker build -t image .

docker-run:
	docker run -p 8080:8080 image

compose:
	docker compose build && docker compose up -d

down:
	docker compose down --remove-orphans

restart: 
	docker compose down --remove-orphans && docker compose build && docker compose up -d

swagger-clean:
	rm -rf generated/restapi/operations/
	rm -f  generated/restapi/doc.go
	rm -f  generated/restapi/embedded_spec.go
	rm -f  generated/restapi/server.go
	rm -f  cmd/main.go

kill-localhost:
	kill $$(lsof -t -i @127.0.0.1:8080 -i @127.0.0.1:8081)

kill-any:
	kill $$(lsof -t -i @0.0.0.0:8080 -i @0.0.0.0:8081)

test:
	@env $(shell cat $(wildcard .env) | xargs)  go test -v -bench -coverpkg=./internal/...,./cmd/... -coverprofile=coverage.cov ./... -json -v | tparse

test-raw:
	@env $(shell cat $(wildcard .env) | xargs)  go test -v -bench -coverpkg=./internal/...,./cmd/... -coverprofile=coverage.cov ./...
