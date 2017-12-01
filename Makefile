GO = go
# PACKAGES is a list of packages to build
PACKAGES = $(shell go list ./...)
# VERSION is the current version number to build
VERSION=`cat .version`
# LDFLAGS are the compiler flags being used
LDFLAGS=-ldflags "-X github.com/tumult-project/tumult/version.Version=${VERSION}"

# Building

install:
	$(GO) install $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:
	@echo $(VERSION)

# Cleaning

clean:
	$(GO) clean
	@rm -rf .coverage.out

# Testing

test:
	$(GO) test -v -coverprofile=.coverage.out -covermode=count $(PACKAGES)

cover: test
	$(GO) tool cover -func=.coverage.out

cover-html: test
	$(GO) tool cover -html=.coverage.out

# Dependencies

deps:

dev-deps:
	$(GO) get -u github.com/alecthomas/gometalinter
	@gometalinter --install 
