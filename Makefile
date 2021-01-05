build: FORCE
	go build -o build/feed-go-mail src/main/main.go
clean:
	rm -rf build
run:
	go run src/main/main.go

FORCE: ;