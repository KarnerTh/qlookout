.DEFAULT_GOAL := build_project

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

.PHONY: publish_package
publish_package:
	GOPROXY=proxy.golang.org go list -m github.com/KarnerTh/qlookout@main
