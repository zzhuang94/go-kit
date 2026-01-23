@echo off
REM Windows batch script to run all tests
echo 🧪 Running all tests...
go test ./...
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Some tests failed!
    exit /b %ERRORLEVEL%
)
echo ✅ All tests passed!

echo.
echo 📊 Running benchmark tests...
go test ./... -bench=. -benchmem
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Some benchmark tests failed!
    exit /b %ERRORLEVEL%
)
echo ✅ Benchmark tests completed!
