APP_NAME:=pgit
VERSION:=1.0.1
LDFLAGS := -X 'partial-git/cmd.Version=$(VERSION)'

.PHONY: build run clean install

build:
	mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) .

run: build
	./bin/$(APP_NAME)

version:
	@echo "Current version: $(VERSION)"

clean:
	rm -rf bin
