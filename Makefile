run:
	go run cmd/main.go
build:
	go build cmd/main.go
run build: 
	go build -o app cmd/main.go && ./app