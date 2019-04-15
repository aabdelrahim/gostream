build-server:
	go build -o server ./backend
	docker build -t gostream:latest .
	rm -f server

start-server:
	docker run --rm --name server -d --mount type=volume,source=audioLibrary,target=/musicLibrary gostream:latest