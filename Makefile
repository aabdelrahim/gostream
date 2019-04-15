build-server:
	go build -o server ./backend
	docker build -t gostream:latest .
	rm -f server

start-server:
	docker run --rm --name server -d --mount type=volume,source=~/docker/volumes/gostream/audioLibrary,target=~/musicLibrary gostream:latest
	# docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 --mount type=volume,source=~/docker/volumes/postgres,target=/postgresql/data  postgres