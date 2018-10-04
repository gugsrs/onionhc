clean:
	rm light
	rm onion_light

build:
	go build -o light light.go
	GOOS=linux GOARCH=mipsle go build -o onion_light light.go

deps:
	go get
