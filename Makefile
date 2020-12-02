DAYS := $(shell find -s src -type d -regex '.*/[0-2][0-9]$$')
GOPATH := ${GOPATH}:$(shell pwd)

.PHONY: build clean run

all: build run

build:
	@mkdir -pv bin
	@env GOPATH=$(GOPATH) go build -o bin $(DAYS:src/%=%)

clean:
	@rm -rvf bin

run:
	@echo running all
	@for day in $(DAYS:src/%=%); do \
		echo; \
		echo $$day; \
		echo i: $$(wc -l input/$${day}.txt); \
		echo o: $$(cat input/$${day}.txt | ./bin/$$day); \
	done
