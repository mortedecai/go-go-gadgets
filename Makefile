GO_TOOL		?= go
REPORTS_DIR ?= .reports
COVERAGE_FILE_NAME ?= coverage
COVERAGE_FILE ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).out
COVERAGE_FLAG ?= -coverprofile=$(COVERAGE_FILE)
TEST_PKG ?= ./...
FORMAT_FLAG ?= -json

TEST_CMD	?= test $(FORMAT_FLAG) $(COVERAGE_FLAG) $(TEST_PKG)
COVER_CMD   ?= tool cover

COVER_TEXT_CMD ?= $(GO_TOOL) $(COVER_CMD) -func=$(COVERAGE_FILE)
COVER_HTML_CMD ?= $(GO_TOOL) $(COVER_CMD) -html=$(COVERAGE_FILE)

HTML_REPORT ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).html
TEXT_REPORT ?= $(REPORTS_DIR)/$(COVERAGE_FILE_NAME).txt

STATIC_CHECK_REPORT_TEXT ?= $(REPORTS_DIR)/static-check.txt
STATIC_CHECK_REPORT_JSON ?= $(REPORTS_DIR)/static-check.json

RESULTS_JSON = $(REPORTS_DIR)/test.json

SOURCE_FILES := $(shell find . -type f -name '*.go')

.PHONY: all
all: coverage staticChecks

.PHONY: full
full: clean all

.PHONY: clean
clean:
	@rm -rf $(REPORTS_DIR)
	@mkdir $(REPORTS_DIR)

.PHONY: setupTools
setupTools:
	@echo "Setting up tools and vendoring dependencies"
	@echo "Tools:"
	@echo "* staticcheck v2022.1.2"
	@go install honnef.co/go/tools/cmd/staticcheck@2022.1.2
	@go mod vendor

.PHONY: coverage
coverage: $(COVERAGE_FILE) htmlCoverage funcCoverage

$(COVERAGE_FILE):
	$(GO_TOOL) $(TEST_CMD) > $(RESULTS_JSON)

.PHONY: staticChecks
staticChecks: staticCheck

.PHONY: staticCheck
staticCheck:
	@staticcheck -f stylish ./... > $(STATIC_CHECK_REPORT_TEXT)
	@staticcheck -f json ./... > $(STATIC_CHECK_REPORT_JSON)

htmlCoverage: $(COVERAGE_FILE)
	$(COVER_HTML_CMD) -o $(HTML_REPORT)

funcCoverage: $(COVERAGE_FILE)
	$(COVER_TEXT_CMD) -o $(TEXT_REPORT)
