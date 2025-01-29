.PHONY: clean init run

all: clean init cmd/main

init:
	go mod tidy
	docker-compose up --build --no-start && docker-compose start

cmd/main: cmd/main.go
	@echo "Building..."
	go build -o $@ $<

clean:
	docker-compose down --volumes