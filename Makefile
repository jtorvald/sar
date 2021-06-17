build:
	GOOS=darwin CGO_ENABLED=0 go build -a -ldflags="-w -s" -trimpath -o ./bin/macos/sar main.go
	GOOS=linux CGO_ENABLED=0 go build -a -ldflags="-w -s" -trimpath -o ./bin/linux/sar main.go

run:
	go run main.go
