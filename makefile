
BINARY_NAME=factor_app

check-deps:
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go не установлен. Установите golang-go."; exit 1; }

build: check-deps
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

deps:
	go mod tidy

.PHONY: build run clean check-deps deps