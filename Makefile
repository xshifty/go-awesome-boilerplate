help:
	@echo "make {help}              show this message"
	@echo "make setup               install requirements to build this app"
	@echo "make templates           generate go files from templ templates"
	@echo "make build               build this app"
	@echo "make run                 initialize dev server"

setup:
	go install github.com/cosmtrek/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	templ generate
	go mod tidy

templates:
	templ generate

build: templates
	go build -o ./build/bootstrap ./cmd/bootstrap/main.go

run: templates
	air

