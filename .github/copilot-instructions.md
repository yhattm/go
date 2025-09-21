# RxGo Testing Library

This is a Go module that provides testing utilities for the RxGo reactive extensions library. The project demonstrates reactive stream patterns using RxGo v2.5.0 including observables, intervals, and event sources.

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Prerequisites and Setup
- Ensure Go 1.20+ is properly installed and working
- **CRITICAL**: The Go installation in some environments may have internal errors. If you encounter "internal error: missing go version" or "missing go root module" errors, this indicates a broken Go installation that prevents normal operations.

### Building and Testing
- **Standard Go Commands**: 
  - `go mod download` -- downloads dependencies. Takes 30-60 seconds. NEVER CANCEL. Set timeout to 2+ minutes.
  - `go mod tidy` -- cleans up dependencies. Takes 10-30 seconds.
  - `go build ./...` -- builds all packages. Takes 10-30 seconds. NEVER CANCEL. Set timeout to 2+ minutes.
  - `go test ./...` -- runs all tests. Takes 30-60 seconds. NEVER CANCEL. Set timeout to 2+ minutes.
  - `go test -v ./...` -- runs tests with verbose output
  - `go test ./lib/rxgotest/` -- runs tests for specific package

### **CRITICAL LIMITATION**: 
- **IF GO COMMANDS FAIL**: If standard Go commands fail with internal errors, document this as "Go installation is broken - standard go commands fail with internal errors". DO NOT attempt workarounds as this is an environment issue, not a code issue.

### Running Code
- Navigate to the library: `cd lib/rxgotest/`
- Run specific examples: `go run rxgotest.go` (if Go installation works)
- Test individual functions through the test files

## Validation

### **MANUAL VALIDATION REQUIREMENTS**:
- **ALWAYS** validate that Go is working before attempting any operations: `go version`
- If Go commands work, ALWAYS run the full test suite: `go test -v ./...`
- **Test Scenario Validation**: After making changes, run specific test scenarios:
  - `go test -v ./lib/rxgotest/ -run Test_just` -- tests basic observable creation
  - `go test -v ./lib/rxgotest/ -run Test_Interval` -- tests interval-based observables  
  - `go test -v ./lib/rxgotest/ -run Test_FromChannel` -- tests event source observables
- **NEVER CANCEL long-running tests** - some reactive tests include deliberate delays (3-5 seconds per test)

### Code Quality
- Run Go formatting: `go fmt ./...`
- Check for potential issues: `go vet ./...`
- **No linting tools configured** - the project does not include additional linters

## Repository Structure

### Key Files and Directories
```
/home/runner/work/go/go/
├── go.mod                    # Go module definition
├── go.sum                    # Dependency checksums
├── .gitignore               # Standard Go gitignore
├── .devcontainer/           # Dev container configuration
│   └── devcontainer.json   # Uses universal dev container
└── lib/
    └── rxgotest/           # Main library code
        ├── rxgotest.go     # Observable creation utilities
        └── rxgotest_test.go # Test cases
```

### **Package: lib/rxgotest**
- **Primary functions** (see `/home/runner/work/go/go/lib/rxgotest/rxgotest.go`):
  - `NewObsFromJust()` - Creates observable from static values (1, 2, 3)
  - `NewObsFromInterval()` - Creates time-based observable that emits every second
  - `NewObsFromEventSource()` - Creates observable from channel events with 3-second timeout
  - `Observe()` - Observes and logs observable emissions, panics on errors
  - `DoOnNext()` - Adds side effects to observable chain with logging
  - `DoOnCompleted()` - Adds completion handler with logging

- **Test functions** (see `/home/runner/work/go/go/lib/rxgotest/rxgotest_test.go`):
  - `Test_just` - Tests basic static observable creation
  - `Test_Interval` - Tests interval observable with Take(5) limitation  
  - `Test_DoOnNext` - Tests observable with context cancellation after 5 seconds
  - `Test_FromChannel` - Tests event source observable with Take(5) and 5-second delay

## Dependencies
- **RxGo v2.5.0** (`github.com/reactivex/rxgo/v2`) - Reactive extensions for Go
- **Go 1.20+** required (specified in go.mod)

## Common Tasks

## Coding Patterns and Best Practices

### Reactive Patterns Used
- **Context cancellation**: Always use `context.WithCancel()` for long-running observables
- **Resource cleanup**: Use `defer cancel()` and `close(ch)` for proper cleanup
- **Error handling**: Check `item.E != nil` in observation loops and panic on errors
- **Observable limiting**: Use `Take(n)` to prevent infinite observables in tests
- **Timing control**: Use `time.After()` for controlled timeouts (3-5 seconds typical)

### Code Style Conventions
- Package name: `rxgotest` (simple, descriptive)
- Function naming: `NewObsFrom*` pattern for observable constructors
- Test naming: `Test_*` with descriptive suffixes
- Import grouping: Standard library first, then external packages
- Error handling: Immediate panic on observable errors (testing context)

### Adding New Observable Patterns
1. Add function to `lib/rxgotest/rxgotest.go`
2. Add corresponding test to `lib/rxgotest/rxgotest_test.go`
3. Follow existing patterns for observable creation and testing
4. Always include context cancellation for long-running observables
5. Use appropriate `Take(n)` limits in tests to prevent infinite execution

### Testing New Features
- Test files use deliberate time delays (3-5 seconds) to test reactive behavior
- Use `Take(n)` to limit infinite observables in tests
- Always include context cancellation for proper cleanup
- Include both immediate observables (Just) and time-based observables (Interval)

### **Important Timing Notes**:
- **Test execution**: Each test includes time.Sleep() calls for reactive testing - this is intentional
- **NEVER CANCEL** test runs as they include deliberate delays up to 5 seconds per test
- **Total test runtime**: Expect 15-30 seconds for full test suite

## Environment Limitations
- **Network access may be restricted** - dependency downloads may fail
- **Go installation may be broken** - if go commands fail with internal errors, this is an environment issue
- **No CI/CD pipeline configured** - no GitHub Actions or automated testing

## Troubleshooting

### Common Issues and Solutions

#### Go Installation Problems
- **Error**: "internal error: missing go version" or "missing go root module"
- **Solution**: Document as environment issue, do not attempt fixes
- **Workaround**: Focus on static code analysis only

#### Network/Dependency Issues  
- **Error**: "could not resolve host" during `go mod download`
- **Solution**: Document as network restriction, note dependencies in go.sum are already resolved
- **Workaround**: Dependencies should already be cached in go.sum

#### Test Timeouts
- **Issue**: Tests appear to hang
- **Explanation**: Tests include deliberate 3-5 second delays for reactive testing
- **Solution**: Wait for completion, NEVER CANCEL
- **Expected**: Full test suite takes 15-30 seconds

#### Observable Infinite Loops
- **Issue**: Interval observables run indefinitely  
- **Solution**: Always use `Take(n)` in tests to limit emissions
- **Pattern**: `obs.Take(5)` is commonly used in existing tests

## Quick Reference Commands

### Essential Commands (when Go works)
```bash
# Verify setup
go version
cd /home/runner/work/go/go

# Build and test
go mod download           # timeout: 120s
go build ./...           # timeout: 120s  
go test -v ./...         # timeout: 120s, expect 15-30s runtime

# Code quality
go fmt ./...
go vet ./...

# Specific tests
go test -v ./lib/rxgotest/ -run Test_just      # ~2s
go test -v ./lib/rxgotest/ -run Test_Interval  # ~10s  
go test -v ./lib/rxgotest/ -run Test_FromChannel # ~8s
```

### When Go is Broken
```bash
# Verify issue
go version  # Should show version or error

# Focus on static analysis  
find . -name "*.go" -exec wc -l {} +
grep -r "Observable\|rxgo" lib/
```

## Reference Information

### Expected Command Outputs (when Go is working correctly)
```bash
# Go version check
$ go version
go version go1.20.x linux/amd64

# Project structure
$ ls -la
total 32
drwxr-xr-x 5 runner runner 4096 Sep 21 12:46 .
drwxr-xr-x 3 runner runner 4096 Sep 21 12:45 ..
drwxrwxr-x 2 runner runner 4096 Sep 21 12:46 .devcontainer
drwxrwxr-x 7 runner runner 4096 Sep 21 12:46 .git
-rw-rw-r-- 1 runner runner  478 Sep 21 12:46 .gitignore
-rw-rw-r-- 1 runner runner  479 Sep 21 12:46 go.mod
-rw-rw-r-- 1 runner runner 3391 Sep 21 12:46 go.sum
drwxrwxr-x 3 runner runner 4096 Sep 21 12:46 lib

# Key Go files
$ find . -name "*.go"
./lib/rxgotest/rxgotest_test.go
./lib/rxgotest/rxgotest.go
```

### go.mod content
```
module go

go 1.20

require github.com/reactivex/rxgo/v2 v2.5.0

require (
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
```

## Working Around Go Installation Issues
**If go commands fail with internal errors**:
1. Document the limitation: "Go installation broken - standard go commands fail with internal errors"
2. Focus on code analysis and static changes only
3. DO NOT attempt environment fixes or alternative Go installations  
4. Note that the code structure and dependencies appear correct based on go.mod analysis
5. **Common error patterns**: "internal error: missing go version" or "missing go root module" indicate broken Go runtime