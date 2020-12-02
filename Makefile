.PHONY: build run clean

all: build run

build:
	@mkdir -pv bin
	@go build -o bin $$(go list ./...)

clean:
	@rm -rvf bin
