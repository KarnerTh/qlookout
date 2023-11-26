.DEFAULT_GOAL := build_project

# GIT_TAG := $(shell git describe --tags --abbrev=0)
GIT_TAG := 0.0.0-$(shell git rev-parse HEAD) # use commit hash until we switch to proper versions

.PHONY: build_project
build_project:
	@echo "Building app .."
	go build -o qlookout main.go
	@echo "Building core done ✅"

.PHONY: build_web
build_web:
	@echo "Building web app.."
	# prefix stands for the folder where to source is
	npm run build --prefix web
	@echo "Building web app done ✅"

.PHONY: publish-package
publish-package:
	GOPROXY=proxy.golang.org go list -m github.com/KarnerTh/query-lookout@$(GIT_TAG)
