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
