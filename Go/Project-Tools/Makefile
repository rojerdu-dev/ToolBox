BINARY_NAME=api

build:
	@GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}

run: build
	@./bin/${BINARY_NAME}

clean:
	@go clean
	@rm bin/${BINARY_NAME}

test:
	@go test -v ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out

vet:
	@go vet

format:
	@gofmt -s -w .
	@goimports -w .
	@echo "\n go fmt and imports -> done!"

lint:
	@golangci-lint run

seed:
	@go run scripts/seed.go

docker:
	echo "building docker file"
	@docker build -t api .
	echo "running API inside Docker container"
	@docker run -p 8000:8000 api
