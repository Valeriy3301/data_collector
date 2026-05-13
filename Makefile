run:
	go run main.go

build:
	go build -o collector .

test:
	go test ./...

docker:
	docker build -t collector .