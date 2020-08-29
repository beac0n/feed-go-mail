deps_install:
	dep ensure
build: FORCE
	go build -o build/feedgomail src/main/main.go
clean:
	rm -rf build
run:
	go run src/main/main.go

FORCE: ;