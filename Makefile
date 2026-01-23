.PHONY: test test-verbose test-coverage clean help

# 默认目标
.DEFAULT_GOAL := help

# 运行所有测试
test:
	@echo "🧪 Running all tests..."
	@go test ./...

# 运行所有测试（详细输出）
test-verbose:
	@echo "🧪 Running all tests with verbose output..."
	@go test ./... -v

# 运行所有测试并显示覆盖率
test-coverage:
	@echo "🧪 Running all tests with coverage..."
	@go test ./... -cover

# 运行所有测试并生成覆盖率报告
test-coverage-html:
	@echo "🧪 Running all tests and generating coverage report..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

# 清理生成的文件
clean:
	@echo "🧹 Cleaning up..."
	@rm -f coverage.out coverage.html
	@find . -name "*.test" -type f -delete
	@echo "✅ Cleanup complete"

# 显示帮助信息
help:
	@echo "go-kit Makefile Commands:"
	@echo ""
	@echo "  make test              - Run all tests"
	@echo "  make test-verbose      - Run all tests with verbose output"
	@echo "  make test-coverage     - Run all tests with coverage"
	@echo "  make test-coverage-html - Run tests and generate HTML coverage report"
	@echo "  make clean             - Clean generated files"
	@echo "  make help              - Show this help message"
	@echo ""
