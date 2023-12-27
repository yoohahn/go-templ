init:
	@yarn install
	@go get github.com/a-h/templ
	@go get github.com/labstack/echo/v4
	@go get github.com/labstack/echo/v4/middleware

templ:
	@templ generate

templWatch:
	@templ generate --watch

run:
	@make dev
	@go run main.go

dev:
	@make templ
	@yarn dev

build:
	@make templ
	@yarn build
	@go build main.go
