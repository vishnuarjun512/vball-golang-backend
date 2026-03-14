# Architecture

This document describes the overall system architecture of the Volleyball Game Backend.

---

## Overview

The backend is a REST API written in Go using the Gin framework. It serves two distinct consumers:

- **Admin Panel** — a web-based dashboard for managing players, abilities, and servers
- **Unreal Engine Game Servers** — dedicated multiplayer servers that query player loadouts and report player connections

All data is persisted in a PostgreSQL database. The entire stack runs inside Docker containers for consistent development and deployment environments.

---

## System Components

### Go Backend (Gin)

The core API server. It handles all incoming HTTP requests, applies business logic, and communicates with the database.

Responsibilities:

- Steam authentication and player registration
- Player loadout assembly and delivery
- Ability CRUD operations
- Matchmaking logic (assigning players to servers)
- Admin management endpoints

### PostgreSQL

The primary data store. It holds all persistent state including players, abilities, servers, bans, and activity logs.

Migrations and seed data are managed through files in the `migrations/` directory.

### Docker

All services are containerized. `docker compose` brings up the Go backend, PostgreSQL, and any supporting services together.

Air is used inside the Go container for hot reload during development.

---

## Request Flow

### Admin Panel Request

```
Admin Browser
    → POST /auth/steam-login        (authenticate)
    → GET  /admin/players           (list players)
    → GET  /admin/players/:steamid  (view player + loadout)
    → GET  /admin/abilities         (view abilities)
```

### Unreal Engine Game Server Request

```
Unreal Server
    → GET /game/player/:steamid/loadout   (fetch player ability IDs before match)
    → GET /game/abilities                 (fetch ability definitions)
    → POST /game/matchmaking/join         (player joins queue)
    → POST /game/matchmaking/leave        (player disconnects)
    → POST /game/matchmaking/sync         (server syncs connected players)
```

Game servers receive **only ability IDs**, not full ability objects, to keep network payloads small.

---

## Matchmaking Flow

When a player presses Play in the Unreal client:

1. Client sends `POST /game/matchmaking/join` with `playerId` and `region`
2. Backend queries `game_servers` for a server in the matching region with available capacity
3. Backend inserts a row into `server_players` and increments `current_players` on the server
4. Backend returns the server's `ip` and `port` to the client
5. Client connects directly to the Unreal server at `ip:port`

When a player disconnects:

1. Unreal server sends `POST /game/matchmaking/leave`
2. Backend removes the player from `server_players` and decrements `current_players`

Servers can also send `POST /game/matchmaking/sync` to reconcile the full player list if the state drifts.

---

## Infrastructure Layout

Physical infrastructure is modelled as a two-level hierarchy:

```
Region (eu / asia / us)
    └── Machine (VPS / physical server with IP)
            └── Game Server (Unreal instance on machine:port)
                    └── Server Players (connected steam IDs)
```

A single machine can run multiple Unreal server instances on different ports. Each game server is uniquely identified by its `machine_id + port` combination, which maps to a public address like `103.21.44.12:7777`.

---

## Project Structure

```
cmd/
    api/            → application entrypoint (main.go)
internal/
    database/       → database connection and pool setup
    handlers/       → HTTP handlers (one file per route group)
    services/       → business logic layer
    repositories/   → database queries (SQL via pgx)
    models/         → data models and structs
    routes/         → route registration and grouping
migrations/         → SQL schema files and seed data
docker/             → Dockerfile and container configuration
docs/               → project documentation
```

### Layer Responsibilities

| Layer      | Package         | Responsibility                                    |
| ---------- | --------------- | ------------------------------------------------- |
| Handler    | `handlers/`     | Parse HTTP request, call service, return response |
| Service    | `services/`     | Business logic, validation, orchestration         |
| Repository | `repositories/` | SQL queries, database access                      |
| Model      | `models/`       | Shared data structures and types                  |

---

## Database Relationships

```
regions
    └── machines (region_id → regions.id)
            └── game_servers (machine_id → machines.id)
                    └── server_players (server_id → game_servers.id)

players
    └── player_abilities (player_id → players.player_id)
            ├── main_abilities (main_ability_id → main_abilities.id)
            └── sub_abilities  (sub_ability_slot1/2/3 → sub_abilities.id)

players
    └── bans (player_id → players.player_id)
```

---

## Technology Choices

| Concern          | Choice        | Reason                                                   |
| ---------------- | ------------- | -------------------------------------------------------- |
| Language         | Go            | Fast, low memory, simple concurrency for game backends   |
| Framework        | Gin           | Lightweight HTTP router with good performance            |
| Database         | PostgreSQL    | Relational integrity for players, abilities, and servers |
| Driver           | pgx           | Performant native PostgreSQL driver for Go               |
| Containerization | Docker        | Consistent environments across dev and production        |
| Hot Reload       | Air           | Fast feedback loop during local development              |
| Game Engine      | Unreal Engine | Target multiplayer client — backend is engine-agnostic   |
| Auth             | Steam         | Native identity for game players                         |
