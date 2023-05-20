build-UserFunction:
	GOOS=linux GOARCH=amd64 go build -o userApi cmd/lambda/main.go
	cp ./userApi $(ARTIFACTS_DIR)/