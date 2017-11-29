GO = go
# PACKAGES is a list of packages to build
PACKAGES = $(shell go list ./...)

# Building

install:
	$(GO) install -v $(PACKAGES)

# Cleaning

clean:
	$(GO) clean
	@rm -rf coverage-all.out

# Testing

test:
	$(GO) test -v $(PACKAGES)

cover-profile:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		$(GO) test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	@rm -rf coverage.out

cover: cover-profile
	$(GO) tool cover -func=coverage-all.out

cover-html: cover-profile
	$(GO) tool cover -html=coverage-all.out

# Dependencies

deps:

dev-deps:
	$(GO) get -u github.com/alecthomas/gometalinter
	@gometalinter --install 
