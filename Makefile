all: server

server:
	(cd src/server && go build && mv server ../../)

clean:
	rm server
