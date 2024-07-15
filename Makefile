APP_NAME = audiop
BIN_DIR = /usr/local/bin

build:
	go build -o $(APP_NAME) cmd/main.go

install: build
	sudo cp $(APP_NAME) $(BIN_DIR)

uninstall:
	sudo rm -f $(BIN_DIR)/$(APP_NAME)

clean:
	rm -f $(APP_NAME)

.PHONY: build install uninstall clean
