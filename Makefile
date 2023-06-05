ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
OUTPUT_DIR := $(CURDIR)/_output

.PHONY: build
build:tidy create_dir
	@go build -o ${OUTPUT_DIR}/apmmonitor -v ./cmd/main.go

.PHONY: run
run: build
	 @${OUTPUT_DIR}/apmmonitor -c ./configs/apm-monitor.yaml

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: create_dir
create_tmp_dir:
	@if [ ! -d "${ROOT_DIR}/_output" ]; then \
		mkdir ${ROOT_DIR}/_output; \
	fi
	@if [ ! -d "${ROOT_DIR}/tmp" ]; then \
		mkdir ${ROOT_DIR}/tmp; \
	fi
