GO_TOOL		?= go
REPORTS_DIR ?= .reports
COVERAGE_FILE_NAME ?= coverage
COVERAGE_FILE ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).out
COVERAGE_FLAG ?= -coverprofile=$(COVERAGE_FILE)
TEST_PKG ?= ./...

TEST_CMD	?= test $(COVERAGE_FLAG) $(TEST_PKG)
COVER_CMD   ?= tool cover

COVER_TEXT_CMD ?= $(GO_TOOL) $(COVER_CMD) -func=$(COVERAGE_FILE)
COVER_HTML_CMD ?= $(GO_TOOL) $(COVER_CMD) -html=$(COVERAGE_FILE)

HTML_REPORT ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).html
TEXT_REPORT ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).coverage

SOURCE_FILES := $(shell find . -type f -name '*.go')

.PHONY: all
all: coverage

.PHONY: full
full: clean coverage

.PHONY: clean
clean:
	@rm -rf $(REPORTS_DIR)
	@mkdir $(REPORTS_DIR)

.PHONY: coverage
coverage: $(COVERAGE_FILE) htmlCoverage funcCoverage

$(COVERAGE_FILE):
	$(GO_TOOL) $(TEST_CMD)

htmlCoverage: $(COVERAGE_FILE)
	$(COVER_HTML_CMD)

funcCoverage: $(COVERAGE_FILE)
	$(COVER_TEXT_CMD)
