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
	go generate ./...

.PHONY: test-coverage
test-coverage:
	go test -cover -coverprofile=coverage.out ./...; \
	cat coverage.out | grep -v "mock" > coverage.clean.out; \
	go tool cover -html=coverage.clean.out -o coverage.html; \
	rm coverage.out coverage.clean.out

# https://github.com/mfridman/tparse is needed
.PHONY: test-all
test-all:
	go test $(test_target) -json | tparse -all

# https://github.com/mfridman/tparse is needed
.PHONY: test-unit
test-unit:
	go test --short $(test_target) -cover -json | tparse -all

.PHONY: test-cleanup
test-cleanup:
	go clean -testcache

.PHONY: merge-graphql-files
merge-graphql-files:
	gqlmerge usecase delivery/graphql/schema.graphql
	@printf '"""\n%s\n"""\n\n%s\n' 'This is autogenerated - use "make merge-graphql-files"' "$$(cat delivery/graphql/schema.graphql)" > delivery/graphql/schema.graphql
	@echo Merge them successfully

.PHONY: prepare-sqlite
prepare-sqlite:
	rm -f qlookout_test.db
	sqlite3 qlookout_test.db < test/db-setup.sql
