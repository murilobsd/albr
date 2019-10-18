lint:
	golint -set_exit_status .

test:
	go test . -v -timeout=30s -parallel=4

cover:
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
	
fmt:
	gofmt -w -s .
