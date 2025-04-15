## redis-tryout

Quick CLI that experiments with Redis functionality in Golang.

### Local Development

To start the application first run `docker compose up` which spins up the Redis instance at
`localhost:6379`. Then invoke `make build` to build the executable.

To interact with Redis, run `./redis-tryout -h`

