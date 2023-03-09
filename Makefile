run:
	@go run $(root_path)/main.go serve

serve:
	@echo "fasten your belts..."
	@$(root_path)/$(app_name) serve

build:
	@cd $(root_path) && go build -o $(app_name) main.go
	@echo "project build successfully!"

start: clean build serve

clean:
	@rm -f  $(root_path)/$(app_name)
	@echo "project cleaned!"

app_name:= "cqrs"
root_path:= ./src