#!/bin/bash
# Unix/Linux/Mac shell script to run all tests

echo "🧪 Running all tests..."
go test ./...

if [ $? -eq 0 ]; then
    echo "✅ All tests passed!"
else
    echo "❌ Some tests failed!"
    exit 1
fi

echo ""
echo "📊 Running benchmark tests..."
go test ./... -bench=. -benchmem

if [ $? -eq 0 ]; then
    echo "✅ Benchmark tests completed!"
else
    echo "❌ Some benchmark tests failed!"
    exit 1
fi
