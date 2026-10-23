# lzk-go-lib

`lzk-go-lib` is a collection of reusable helpers for Go applications, including
environment configuration, database connection setup, application errors,
database type constants, and structured logging.

## Installation

```sh
go get github.com/LzKelvinSoares/lzk-go-lib
```

## Packages

### `environment`

Read server settings from `SERVER_ADDRESS` and `SERVER_PORT`:

```go
package main

import (
	"fmt"

	"github.com/LzKelvinSoares/lzk-go-lib/environment"
)

func main() {
	server := environment.GetServerVariables()
	fmt.Println(server.GetAddressToServe())
}
```

`GetServerVariables` calls `godotenv.Load()` and terminates the process if
either server variable is empty. The `.env` file is optional if the variables
are already present in the process environment.

Database settings are read by `GetDBVariables` from:

| Variable | Description |
| --- | --- |
| `DB_USER` | Database username |
| `DB_PASSWORD` | Database password |
| `DB_ADDRESS` | Database host |
| `DB_PORT` | Database port |
| `DB_NAME` | Database name |

`DBVariables` provides data source helpers and the package provides
`GetSQLClient` and `GetNoSQLClient`. `GetSQLClient` relies on a compatible
`database/sql` driver being registered by the application; this module does
not register database drivers. The helpers do not validate that all database
variables are set.

Database types are defined in `constants`: `MySQL`, `SQLServer`, `MongoDB`,
and `CosmosDB`. At present, `DBVariables.GetDataSource` handles `MySQL`,
`SQLServer`, and `MongoDB`; it does not handle `CosmosDB`.

### `errs`

`errs.AppError` represents an application error with a message and an optional
HTTP status code. Constructors are provided for not found, unexpected,
validation, authentication, and authorization errors:

```go
appErr := errs.NewValidationError("Email address is invalid")
```

`AppError.AsMessage()` returns a copy containing only the message.

### `logger`

The `logger` package exposes `Info`, `Debug`, and `Error` functions. Each
accepts a message and optional `zap.Field` values:

```go
logger.Info("request completed", zap.String("method", "GET"))
```

The logger uses Zap's production configuration and ISO 8601 timestamps.

### `constants`

The `constants` package defines the `DBType` string type and database type
values used by the environment helpers.
