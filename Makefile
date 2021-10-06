build:
	GO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o mirror-status .