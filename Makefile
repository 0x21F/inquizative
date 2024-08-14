build:
	@templ generate .
	@tailwindcss -i views/app.css -o static/styles.css
	@go build -o bin/fullstackgo main.go 

test:
	@go test -v ./...
	
run: build
	@./bin/fullstackgo

tailwind:
	@./tailwindcss -i views/app.css -o static/styles.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:8080
