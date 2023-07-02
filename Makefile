run:
	go run main.go
build:
	go build -o dist/game.exe
dep:
	go mod tidy
	
