run: build 
	@echo "Building the project..."
	./bin/goredis

build:
	go build -o ./bin/goredis .



