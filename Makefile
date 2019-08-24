test:
	APP_ENV=testing CONFIG_DIR=${PWD} go test -v -cover -count 1 -failfast ./...

tests: test
