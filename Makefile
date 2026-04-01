.PHONY: run build clean

run:
	@echo Starting the Go server...	
	go run ./cmd/api

build:
	go build -o api.exe ./cmd/api
	@echo Build complete

clean:
	@echo Cleaning up old binaries...
	rm -f api.exe
	@echo Done