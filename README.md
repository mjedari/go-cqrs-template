# Go CQRS Template

CQRS + Clean Architecture sample in Go with Redis storage and a simple event bus.
This repo is intentionally structured by layer (tech) so new contributors can scan the stack quickly.

## What you get
- CQRS split between command handlers and query handlers.
- Thin HTTP layer with Gorilla Mux controllers.
- Redis-backed repositories plus a Redis queue for events.
- Wiring layer that assembles dependencies without a DI framework.

## Request flow (high level)
1) HTTP request enters `web/controller`.
2) Command/query handler in `app/<feature>` executes use-case logic.
3) Repository reads/writes entities using `infra/providers/storage`.
4) Optional events published to the event bus, queued in Redis, and handled by workers.

## Project structure
```
config/
  config.yaml.sample
src/
  cmd/                 # Cobra commands (serve, config boot)
  web/
    controller/        # HTTP handlers
    route/             # Router + routes
    middleware/        # HTTP middleware (logging)
    wiring/            # Dependency wiring + event bus loop
    config/            # Config structs
  app/
    <feature>/         # Commands, queries, handlers, repositories
    providers/         # App-facing interfaces (storage, messaging)
  domain/
    <feature>/         # Entities, value objects, events
    contract/          # IEntity contract for storage
  infra/
    providers/         # Redis storage + event bus + queue
  common/              # Shared message structs
```

## Layer walkthrough
### `web/` (HTTP boundary)
Routes live in `src/web/route/route.go`. Controllers call command/query handlers and serialize responses.

### `app/` (use cases)
Each feature (user, coin, order) has:
- Commands and command handlers for writes.
- Queries and query handlers for reads.
- Repositories that use the storage interface.

Example: `src/app/order/command_handler.go` validates balance, creates an order, and publishes events.

### `domain/` (core model)
Pure entities and domain events:
- `src/domain/user/entity.go`
- `src/domain/coin/entity.go`
- `src/domain/order/entity.go`
- `src/domain/order/event.go`

### `infra/` (providers)
Redis storage and the event bus live here:
- `src/infra/providers/storage/redis.go` implements the storage interface.
- `src/infra/providers/messaging/*` implements a Redis-backed queue and event bus.

### `web/wiring/` (composition root)
`src/web/wiring/wiring.go` builds the event bus, registers handlers, and runs a worker loop.

## Configuration
Copy `config/config.yaml.sample` to `config/config.yaml` and adjust host/ports:
```yaml
Server:
  Host: localhost
  Port: 8080
Redis:
  Host: localhost
  Port: 6379
```

## Run locally
### Docker
```
docker compose up
```

### Make
Ensure Redis is running and Go is installed:
```
make start
```

## API quickstart
### create user
```
POST http://localhost:8080/user/create
{
  "name": "Hadi",
  "balance": 1000
}
```

### get user(s)
```
GET http://localhost:8080/user/{user_id}
GET http://localhost:8080/user/all
```

### create coin
```
POST http://localhost:8080/coin/create
{
  "name": "ABAN",
  "price": 3,
  "min": 10
}
```

### get coin(s)
```
GET http://localhost:8080/coin/{coin_id}
GET http://localhost:8080/coin/all
```

### create order
```
POST http://localhost:8080/order/create
{
  "user_id": 1,
  "coin_id": 1,
  "quantity": 2
}
```

### get order(s)
```
GET http://localhost:8080/order/all
```

## For AI advocates (vibe coding)
If you want to use AI as your coding companion during early stage development (vibe coding), using this
go-cqrs-template as the architectural foundation gives the assistant a structured context. That way, AI
suggestions focus on building features that fit into a solid pattern instead of guessing how to organize
modules or responsibilities.

This is especially useful if someone wants to start a Go application by giving an AI a prompt (or even
just linking this repo). The assistant can copy the structure and architecture directly instead of
inventing new layers or folder layouts.

How to use this repo with AI tools:
1) Give the AI the repo URL (or a zip) and ask it to follow the structure.
2) Tell it to keep entities in `src/domain`, use cases in `src/app`, HTTP in `src/web`, and providers in `src/infra`.
3) Ask for a feature and the exact layer changes you want (commands/queries, handlers, routes, wiring).

Example prompt:
```
Use https://github.com/mjedari/go-cqrs-template as the base structure.
Create a "wallet" feature with commands/queries in src/app/wallet,
entities in src/domain/wallet, HTTP handlers in src/web/controller/wallet.go,
and wire it in src/web/wiring and src/web/route/route.go.
Do not introduce new frameworks.
```

## License
MIT. See `LICENSE`.

## Contributing
Contributors are welcome. Open an issue with your idea or submit a PR with a clear description of the change.
