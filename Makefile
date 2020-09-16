SHELL=/bin/bash
BIN=alfred-workflow-search-rubygems
PACKAGE=alfred-workflow-search-rubygems.alfredworkflow
DIST=.
GO=go
export GO111MODULE=on

.PHONY: all
all: build

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 ${GO} build -o ${DIST}/${BIN} .

.PHONY: pack
pack: build
	zip -r ${DIST}/${PACKAGE} ${DIST}/${BIN} info.plist icon.png

.PHONY: clean
clean:
	rm -f ${DIST}/${BIN}
	rm -f ${DIST}/${PACKAGE}
	${GO} clean
