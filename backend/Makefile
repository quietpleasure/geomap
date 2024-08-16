.PHONY: build
build:
	go build -ldflags "-w -s" -v ./cmd/geomap
# -ldflags "-w -s"

.PHONY: tidy
tidy:
	go mod tidy

my_date = $(shell date /t)
.PHONY: push
push:
	git a 
	git co "${my_date}"
	git pusm

.PHONY: pull
pull:
	git pull origin master

.PHONY: test
test:
	go test -v ./internal/app/...
.DEFAULT_GOAL := build

.PHONY: front
front:
	serve -s frontend/dist