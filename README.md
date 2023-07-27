# logger - Simple Logging Library for Go (Golang)

[![Go Report Card](https://goreportcard.com/badge/github.com/deryl-sagala/logger)](https://goreportcard.com/report/github.com/deryl-sagala/logger)[![Go Reference](https://pkg.go.dev/badge/github.com/deryl-sagala/logger.svg)](https://pkg.go.dev/github.com/deryl-sagala/logger)

`logger` is a lightweight logging library for Go (Golang) that provides simple APIs for logging messages at different levels. The API is inspired by Ruby's `logger.rb` and JavaScript's `console.log()` methods. It allows developers to log messages at various severity levels and choose between logging to the console (`os.Stdout`) or directly to a file.

## Features

- Five different log levels: 'fatal', 'error', 'warn', 'info', 'debug'
- Set the maximum log level at runtime to control which messages are logged
- Choose to log to the console or a file
- Simple and intuitive API

## Installation

```bash
go get github.com/deryl-sagala/logger
```

## Usage

Here's a quick guide on how to use the `logger` library:

1. Create a new logger instance using `logger.NewLogger()`.

```go
import "github.com/deryl-sagala/logger"

log := logger.NewLogger()
```

2. (Optional) Set the maximum log level. Only messages with the specified level and lower will be logged.

```go
log.SetMaxLogLevel(logger.Debug) // Set the maximum log level to Debug
```

3. (Optional) If you want to log directly to a file, set the log file path using `SetLogFile()`.

```go
err := log.SetLogFile("logs.txt")
if err != nil {
    // Handle error if unable to open the log file
}
```

4. Use the appropriate log level methods to log messages at different levels.

```go
log.Fatal("This is a fatal error.")
log.Error("This is an error message.")
log.Warn("This is a warning.")
log.Info("This is an info message.")
log.Debug("This is a debug message.")
```

## Log Levels

- `Fatal`: Logs a fatal-level message and exits the application.
- `Error`: Logs an error-level message.
- `Warn`: Logs a warning.
- `Info`: Logs an informational message.
- `Debug`: Logs a debug-level message.

## Contributions

Contributions to the `logger` library are welcome! If you encounter any issues or have suggestions for improvements, feel free to create an issue or submit a pull request.

## License

This library is distributed under the MIT License. See the [LICENSE](LICENSE) file for more information.

Happy Logging with `logger` in Go (Golang)!
```
