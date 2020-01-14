GO = go
GOM = GO111MODULE=on GOFLAGS=-mod=vendor $(GO)
MOCKGEN ?= mockgen
VERSION = 0.1.0
LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse HEAD) -X main.version=$(VERSION) -X main.name=$(BINARY)"

BINDIR = bin
BINARY = hypotheses

OS := $(shell uname)

# building vars
ifeq ($(OS),Linux)
	TMPDIR = /tmp
	REPORTDIR = /opt/atlassian/pipelines/agent/build/test-reports
endif
ifeq ($(OS),Darwin)
	REPORTDIR = $(TMPDIR)test-reports
endif

$(BINDIR)/$(BINARY): $(BINDIR)
ifeq ($(OS),Darwin)
	GOOS=darwin $(GOM) build -buildmode=pie -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)
endif
ifeq ($(OS),Linux)
	GOOS=linux $(GO) build -buildmode=pie -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)
endif

$(REPORTDIR):
	mkdir -p $(REPORTDIR)

$(BINDIR):
	mkdir -p $(BINDIR)

build: $(BINDIR)/$(BINARY)

.PHONY: deps
deps:
	$(GOM) mod tidy
	$(GOM) mod vendor

.PHONY: test
test: $(REPORTDIR)
	$(GOM) test -buildmode=pie -v -cover ./... 2>&1 | tee $(TMPDIR)results.txt
	cat $(TMPDIR)results.txt | go-junit-report -set-exit-code > $(REPORTDIR)/gotest.xml

.PHONY: check
check: $(REPORTDIR)
	if [ -d vendor ]; then cp -r vendor/* ${GOPATH}/src/; fi
	GO111MODULE=off $(GO) get -v github.com/securego/gosec/cmd/gosec/...
	GO111MODULE=off gosec -exclude G304 ./...

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINDIR)/*


