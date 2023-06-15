.DEFAULT_GOAL := build

.PHONY: build
build: build_web build_core

.PHONY: build_web
build_web:
	@echo "Building web app.."
	# prefix stands for the folder where to source is
	npm run build --prefix web
	@echo "Building web app done ✅"

.PHONY: build_core
build_core:
	@echo "Building core .."
	CGO_ENABLED=1 go build -o qlookout main.go
	@echo "Building core done ✅"
