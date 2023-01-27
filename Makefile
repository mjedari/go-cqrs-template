run:
	@go run main.go serve

serve:
	@echo "fasten your belts..."
	@./$(app_name) serve

build:
	@go build -o $(app_name) main.go
	@echo "project build successfully!"

start: clean build serve

clean:
	@rm -f $(app_name)
	@echo "project cleaned!"

app_name:= "cqrs"