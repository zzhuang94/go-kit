@echo off
REM Windows batch script to run all tests
echo 🧪 Running all tests...
go test ./...
if %ERRORLEVEL% EQU 0 (
    echo ✅ All tests passed!
) else (
    echo ❌ Some tests failed!
    exit /b %ERRORLEVEL%
)
