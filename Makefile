BINARY_NAME=dist/YuGiGo

build:
	go build -o $(BINARY_NAME) -v

debug: 
	export DEBUG=true
	$(MAKE) run


run: build
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)