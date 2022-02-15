


build:
	rm -rf build
	mkdir build
	go build cmd/proxy/main.go
	mv ./main build/


run:
	go run ./...


docker-build:
	sudo docker build -t proxy .

docker-run:
	sudo docker run -p 8080:8080 -t proxy