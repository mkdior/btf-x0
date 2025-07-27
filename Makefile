.PHONY: test run run-requests \
        request-user-create \
        request-merkle-build \
        request-merkle-proof-generate

test:
	go test ./pkg/...
	@echo "\n"

run:
	go build -o ./bin/server ./cmd/server && ./bin/server
	@echo "\n"

request-user-create:
	curl -X POST http://127.0.0.1:8082/user/create \
		-H 'Content-Type: application/json' \
		-d '{"data":[ \
			{"id":1,"balance":1111}, \
			{"id":2,"balance":2222}, \
			{"id":3,"balance":3333}, \
			{"id":4,"balance":4444}, \
			{"id":5,"balance":5555}, \
			{"id":6,"balance":6666}, \
			{"id":7,"balance":7777}, \
			{"id":8,"balance":8888} \
		]}'
	@echo "\n"

request-merkle-build:
	curl -X POST http://127.0.0.1:8082/merkle/build
	@echo "\n"

request-merkle-proof-generate:
	curl -X GET http://127.0.0.1:8082/merkle/proof/generate \
		-H 'Content-Type: application/json' \
		-d '{"id": 7}'
	@echo "\n"

run-requests: request-user-create request-merkle-build request-merkle-proof-generate
