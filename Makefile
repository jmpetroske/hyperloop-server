all: server

server: src/server/*.go
	(cd src/server && go build && mv server ../../)

clean:
	rm -f server
