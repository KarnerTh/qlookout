.PHONY: create-migration
create-migration:
ifndef name
	$(error name is not set)
endif
	migrate create -ext sql -dir migration $(name)

.PHONY: run-migration
run-migration:
	migrate -path migration -database sqlite3://data.db up

.PHONY: run-down-migration
run-down-migration:
	migrate -path migration -database sqlite3://data.db down 1
