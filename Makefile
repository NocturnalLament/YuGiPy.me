BINARY_NAME=dist/YuGiGo

build:
	clear
	go build -o $(BINARY_NAME) -v

debug:
	clear
	export DEBUG=true
	$(MAKE) run


run: build
	clear
	./$(BINARY_NAME)

clean:
	go clean
	clear
	rm -f $(BINARY_NAME)

.PHONY: build run clean debug