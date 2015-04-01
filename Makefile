#!/usr/bin/make -f
# -*- makefile -*-

# Copyright (C) 2015 Kouhei Maeda <mkouhei@palmtb.net>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.	 See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.	 If not, see <http://www.gnu.org/licenses/>.


BIN := komame
SRC := *.go
GOPKG := github.com/mkouhei/komame
GOPATH := $(CURDIR)/_build
export GOPATH
PATH := $(CURDIR)/_build/bin:$(PATH)
export PATH
# "FLAGS=" when no update package
FLAGS := $(shell test -d $(GOPATH) && echo "-u")

ifdef FLAGS
VENVFLAG := --clear
PIPFLAG := -U
else
VENVFLAG :=
PIPFLAG :=
endif

# "FUNC=-html" when generate HTML coverage report
FUNC := -func

all: precheck clean test build build-docs

precheck:
	@if [ -d .git ]; then \
	set -e; \
	diff -u .git/hooks/pre-commit utils/pre-commit.txt ;\
	[ -x .git/hooks/pre-commit ] ;\
	fi

prebuild:
	go get -d -v ./...
	install -d $(CURDIR)/_build/src/$(GOPKG)
	cp -a $(CURDIR)/*.go $(CURDIR)/testdata $(CURDIR)/_build/src/$(GOPKG)


build: prebuild
	go build -ldflags "-X main.ver $(shell git describe --always)" -o _build/$(BIN)

build-only:
	go build -ldflags "-X main.ver $(shell git describe --always)" -o _build/$(BIN)

prebuild-docs:
	virtualenv $(VENVFLAG) _build/venv
	_build/venv/bin/pip install $(PIPFLAG) -r docs/requirements.txt

build-docs: prebuild-docs
	. _build/venv/bin/activate;\
	cd docs; \
	make html; \
	deactivate

clean:
	@rm -rf _build/$(BIN) $(GOPATH)/src/$(GOPKG)

test: prebuild
	go get $(FLAGS) golang.org/x/tools/cmd/goimports
	go get $(FLAGS) github.com/golang/lint/golint
	go get $(FLAGS) golang.org/x/tools/cmd/vet
	go get $(FLAGS) golang.org/x/tools/cmd/cover
	_build/bin/golint
	go vet
	go test -v -covermode=count -coverprofile=c.out $(GOPKG)
	@if [ -f c.out ]; then \
		go tool cover $(FUNC)=c.out; \
		unlink c.out; \
	fi
	rm -f $(BIN).test main.test
	for src in $(SRC); do \
		gofmt -w $$src ;\
		goimports -w $$src; \
	done
