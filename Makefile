.PHONY: lint test format 

golangciLintVersion = "v1.55.2"
gofumptVersion = "v0.5.0"
gciVersion = "v0.11.0"
govulncheckVersion = "v1.0.1"

# Lint Go Code
$(GOBIN)/golangci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@${golangciLintVersion}

lint: | $(GOBIN)/golangci-lint
	@echo Linting...
	@golangci-lint  -v --concurrency=3 --config=.golangci.yml --issues-exit-code=1 run \
	--out-format=colored-line-number 

# Format Go Code
$(GOBIN)/gofumpt:
	@go install mvdan.cc/gofumpt@${gofumptVersion}
	@go mod tidy

gofumpt: | $(GOBIN)/gofumpt
	@gofumpt -w $(shell ls  -d $(PWD)/*/)

# Format imports
$(GOBIN)/gci:
	@go install github.com/daixiang0/gci@${gciVersion}
	@go mod tidy

gci: | $(GOBIN)/gci
	@gci write --section Standard --section Default --section "Prefix(github.com/wimspaargaren/prolayout)" $(shell ls  -d $(PWD)/*)
	
# Run unit tests and generate coverage report
test:
	@mkdir -p reports
	@go test -coverprofile=reports/codecoverage_all.cov ./... -cover -race -p=4
	@go tool cover -func=reports/codecoverage_all.cov > reports/functioncoverage.out
	@go tool cover -html=reports/codecoverage_all.cov -o reports/coverage.html
	@echo "View report at $(PWD)/reports/coverage.html"
	@tail -n 1 reports/functioncoverage.out

# Opens coverage report in browser
coverage-report:
	@open ./reports/coverage.html

# Formats code
format:
	@make gofumpt
	@make gci
	@go mod tidy

# Install and run govulncheck tool
$(GOBIN)/govulncheck:
	@go install golang.org/x/vuln/cmd/govulncheck@${govulncheckVersion}

govulncheck: | $(GOBIN)/govulncheck
	@govulncheck -test ./...
