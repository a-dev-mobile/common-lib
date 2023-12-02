# Common-Lib

This Go module, `github.com/a-dev-mobile/common-lib`.

## Installation

To use these libraries in your Go project, run:

```shell
go get github.com/a-dev-mobile/common-lib
```

Then, import the needed packages in your Go files.

### Package: config

The `config`` package provides functionality to load and parse YAML configuration files.

#### Usage

Import the package:

```shell
import "github.com/a-dev-mobile/common-lib/config"
```

```go
var cfg *models.Config
var err error

cfg, err = config.GetConfig[models.Config]("../config","config.yaml")
 if err != nil {
  log.Fatalf("Error loading config: %s", err)
 }
    
```

### Package: logging

The `logging`` package provides structured logging with different levels and output options.

#### Usage

Import the package:

```shell
import "github.com/a-dev-mobile/common-lib/logging"
```

Set up the logger:

```go
logLevel := "info" // Can be debug, info, warn, error
logFilePath := "path/to/logfile.log" // Leave empty for stdout
logger := logging.SetupLogger(logLevel, logFilePath)
```

Use the logger:

```go
logger.Info("Starting application")
if err := someFunction(); err != nil {
    logger.Error("Something went wrong", logging.Err(err))
}

```
