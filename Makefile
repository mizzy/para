NAME := para
VERSION = $(shell gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)

.PHONY: build
build:
	go build -o bin/$(NAME)


.PHONY: package
package:
	@sh -c "'$(CURDIR)/scripts/package.sh'"

.PHONY: crossbuild
crossbuild:
	goxz -pv=v${VERSION} -build-ldflags="-X main.GitCommit=${REVISION}" \
	-arch=amd64 -d=./pkg/dist/v${VERSION} \
	-n ${NAME}

.PHONY: release
release: package
	ghr -u mizzy v${VERSION} ./pkg/dist/v${VERSION}

.PHONY: bump
bump:
	@sh -c "'$(CURDIR)/scripts/bump.sh'"
