test_target := "./..."

.PHONY: create-migration
create-migration:
ifndef name
	$(error name is not set)
endif
	migrate create -ext sql -dir migration $(name)

.PHONY: run-migration
run-migration:
	migrate -path database/migrations -database sqlite3://data.db up

.PHONY: run-down-migration
run-down-migration:
	migrate -path database/migrations -database sqlite3://data.db down 1

# https://github.com/vektra/mockery is needed
.PHONY: create-mocks
create-mocks:
	mockery --all

.PHONY: test-coverage
test-coverage:
	go test -cover -coverprofile=coverage.out ./...; go tool cover -html=coverage.out -o coverage.html; rm coverage.out

# https://github.com/mfridman/tparse is needed
.PHONY: test-all
test-all:
	go test $(test_target) -cover -json | tparse -all

# https://github.com/mfridman/tparse is needed
.PHONY: test-unit
test-unit:
	go test --short $(test_target) -cover -json | tparse -all

.PHONY: test-cleanup
test-cleanup:
	go clean -testcache

