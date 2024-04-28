BINARY_NAME=GoldWatcher
APP_NAME=GoldWatcher
VERSION=1.0.1
BUILD_NO=2

build:
	rm -rf bin/
	fyne package -appVersion ${VERSIO} -appBuild ${BUILD_NO} -name ${APP_NAME} --src ./cmd/goldwatcher/main.go -release

run:
	env DB_PATH="./bin/sql.db" go run ./cmd/goldwatcher/main.go

clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf bin/
	@echo "Cleaned!"

test:
	go test -v ./...
