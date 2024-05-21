run-example-async:
	go run cmd/example_async/main.go

run-example-sync:
	go run cmd/example_sync/main.go	

test:
	APP_ENV=testing CONFIG_DIR=${PWD} go test -v -cover -count 1 -failfast ./...

badge:
	gopherbadger > /dev/null

tests: test
