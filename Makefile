deps_install_dev:
	go get github.com/oxequa/realize
deps_install:
	dep ensure
dev:
	realize start
build:
	go build -o build/gofeedtomail-linux-amd64 src/main/main.go
clean:
	rm -rf build
	rm -f coverage.html
run:
	go run src/main/main.go
cov:
	go test -tags=test -coverpkg=./... -cover -coverprofile coverage.html -v ./...
	go tool cover -html=coverage.html
test:
	go test -tags=test -v ./...
