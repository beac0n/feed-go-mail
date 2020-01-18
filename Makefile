deps_install:
	dep ensure
build:
	go build -o build/gofeedtomail src/main/main.go
clean:
	rm -rf build
run:
	go run src/main/main.go
