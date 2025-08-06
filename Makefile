APP_NAME:=pgit
VERSION:=0.0.1
LDFLAGS := -X 'main.Version=$(VERSION)'

.PHONY: build run clean

build:
	mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) .

run: build
	./bin/$(APP_NAME)

clean:
	rm -rf bin
