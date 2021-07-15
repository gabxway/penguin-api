APP = penguin-app
BIN_LINUX = bin/$(APP)
CMD_SRC = cmd/$(APP)/main.go

build: 
	go build -o $(BIN_LINUX) $(CMD_SRC)
run:
	./bin/$(APP)

clean:
	rm -rf bin/