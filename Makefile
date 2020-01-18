deps_install:
	dep ensure
build:
	go build -o build/gofeedtomail-linux-amd64 src/main/main.go
clean:
	rm -rf build
run:
	go run src/main/main.go
