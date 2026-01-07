DIST_DIR := dist
BINARY_NAME := bootstrap
ZIP_NAME := bootstrap.zip

.PHONY: build clean

build: clean
	mkdir -p $(DIST_DIR)
	
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME) cmd/api/main.go
	
	cd $(DIST_DIR) && zip $(ZIP_NAME) $(BINARY_NAME)

clean:
	rm -rf $(DIST_DIR)
