Configuration files which could be used in go application. 

It will use external `github.com/joho/godotenv` package to load configuration from the .env file

Env file variables:

```.dotenv

# state.go variables:
APP_ENV=dev # dev, default: prod
APP_DEBUG=true # true, false
APP_DEBUG_LEVEL=debug # debug, info, warn, error, fatal, panic

# grpc.go variables:
GRPC_HOST=localhost
GRPC_PORT=50050

# http.go variables:
HTTP_HOST=localhost
HTTP_PORT=8080

# pg variables:
PG_DSN="host=localhost port=54321 dbname=note user=note-user password=note-password sslmode=disable"
# PG_DSN=postgres://pguser:password@localhost:5432/db-name?sslmode=disable

# swagger.go variables:
SWAGGER_HOST=localhost
SWAGGER_PORT=8090

```

Example usage:
```go
package main

import "config"

type serviceProvider struct {
  pgConfig      config.PGConfig
}


func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

```
