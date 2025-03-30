# Define environment variables here if needed
ENV_VARS = \
	export PORT_APP= && \
	export OSS_ACCESS_KEY_SECRET= && \
	export OSS_ACCESS_KEY_ID= && \
	export OSS_BUCKET= && \
	export OSS_REGION= && \
	export OSS_ENDPOINT= && \
	export OSS_ACCESS_EXPIRED_URL_TIME= && \
	export OSS_PATH=

# Rule to run the Go application
run:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && $(ENV_VARS) && go run main.go; \
	else \
		$(ENV_VARS) && go run main.go; \
	fi

# Optional: Rule to build the Go application
build:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && $(ENV_VARS) && go build -o myapp main.go; \
	else \
		$(ENV_VARS) && go build -o myapp main.go; \
	fi

# run docker
up:
	docker compose up
	
# Optional: Rule to clean the build artifacts
clean:
	rm -f myapp

watch:
	CompileDaemon --build="go build -o main main.go" --command="./main"
.PHONY: watch

proto-path:
	export PATH=$PATH:$(go env GOPATH)/bin