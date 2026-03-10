# Volleyball Game Backend

A Go-based backend service for a multiplayer volleyball game.

This backend manages:

- player accounts
- player abilities
- game servers
- matchmaking infrastructure
- admin dashboard APIs

It integrates with **Unreal Engine dedicated servers** and a **web-based admin dashboard**.

---

# Tech Stack

Backend language

- Go (Golang)

Web framework

- Gin

Database

- PostgreSQL

Database driver

- pgx / pgxpool

Infrastructure

- Docker
- Docker Compose

Development tooling

- Air (hot reload)

---

# Project Structure

vball-go-backend
│
├── cmd/
│ └── api/
│ └── main.go
│
├── internal/
│ ├── database/
│ ├── models/
│ ├── handlers/
│ ├── services/
│ ├── repositories/
│ └── routes/
│
├── migrations/
│
├── docker/
│
├── docker-compose.yml
├── go.mod
└── README.md

---

# Backend Architecture

The backend uses a layered architecture.
HTTP Request
↓
Routes
↓
Handlers
↓
Services
↓
Repositories
↓
PostgreSQL

---

### Layer Responsibilities

| Layer        | Responsibility        |
| ------------ | --------------------- |
| Routes       | Defines API endpoints |
| Handlers     | Handles HTTP requests |
| Services     | Business logic        |
| Repositories | Database queries      |
| Database     | Persistent storage    |

---

# Ability System Design

Abilities are defined separately from players.

Tables:

- `main_abilities`
- `sub_abilities`

Players reference abilities through:

- `player_abilities`

This design allows:

- ability balancing without modifying players
- smaller API responses
- faster queries
- reusable ability definitions

---

# Data Strategy

Two types of systems use this backend.

## Admin Dashboard

Receives **full ability definitions**.

Example endpoints:

GET /admin/abilities
GET /admin/players

The dashboard maps ability IDs to full ability data.

---

## Game Servers

Game servers receive **only ability IDs**.

Example response:

mainAbilityId
subAbilityIds

This keeps gameplay network payloads small.

---

# API Routes

## Authentication

POST /auth/steam-login

Registers a new player automatically if they do not exist.

---

## Player Administration

GET /admin/players
GET /admin/players/:steamid

Returns player data including ability IDs.

---

## Ability Routes

GET /admin/abilities
GET /admin/main-abilities
GET /admin/sub-abilities

Returns ability definitions used by the admin dashboard.

---

## Game Server Routes

GET /game/player/:steamid/loadout
GET /game/abilities

These endpoints are intended for Unreal Engine game servers.

---

# Database Schema

## players

| Column         | Type          |
| -------------- | ------------- |
| player_id      | UUID (PK)     |
| steam_id       | TEXT (UNIQUE) |
| username       | TEXT          |
| kash           | INTEGER       |
| ban_status     | TEXT          |
| matches_played | INTEGER       |
| wins           | INTEGER       |
| losses         | INTEGER       |
| last_login     | TIMESTAMP     |
| created_at     | TIMESTAMP     |

---

## player_abilities

| Column            | Type                          |
| ----------------- | ----------------------------- |
| player_id         | UUID (PK, FK → players)       |
| main_ability_id   | INTEGER (FK → main_abilities) |
| sub_ability_slot1 | INTEGER (FK → sub_abilities)  |
| sub_ability_slot2 | INTEGER (FK → sub_abilities)  |
| sub_ability_slot3 | INTEGER (FK → sub_abilities)  |

---

## main_abilities

| Column              | Type          |
| ------------------- | ------------- |
| id                  | SERIAL (PK)   |
| name                | TEXT (UNIQUE) |
| description         | TEXT          |
| ability_type        | TEXT          |
| tier                | TEXT          |
| duration            | FLOAT         |
| cooldown            | FLOAT         |
| spike_modifier      | FLOAT         |
| jump_modifier       | FLOAT         |
| set_modifier        | FLOAT         |
| receive_modifier    | FLOAT         |
| ball_force_modifier | FLOAT         |
| created_at          | TIMESTAMP     |

---

## sub_abilities

| Column         | Type          |
| -------------- | ------------- |
| id             | SERIAL (PK)   |
| name           | TEXT (UNIQUE) |
| description    | TEXT          |
| tier           | TEXT          |
| modifier_type  | TEXT          |
| modifier_value | FLOAT         |
| created_at     | TIMESTAMP     |

---

## servers

| Column          | Type      |
| --------------- | --------- |
| id              | TEXT (PK) |
| region          | TEXT      |
| max_players     | INTEGER   |
| current_players | INTEGER   |
| game_mode       | TEXT      |
| status          | TEXT      |
| uptime          | TIMESTAMP |

---

## bans

| Column    | Type                |
| --------- | ------------------- |
| id        | UUID (PK)           |
| player_id | UUID (FK → players) |
| reason    | TEXT                |
| admin_id  | TEXT                |
| ban_start | TIMESTAMP           |
| ban_end   | TIMESTAMP           |
| type      | TEXT                |

---

## activity_logs

| Column     | Type      |
| ---------- | --------- |
| id         | UUID (PK) |
| type       | TEXT      |
| message    | TEXT      |
| created_at | TIMESTAMP |

# Development

Start the backend:

docker compose up --build

Stop the backend:

docker compose down

Reset database:

docker compose down -v
docker compose up --build

---

# Planned Features

Future systems include:

- matchmaking
- server orchestration
- ranked ladder
- seasonal progression
- ability reroll mechanics
- admin moderation tools

---

# System Architecture

Admin Dashboard
│
▼
Go API
│
▼
PostgreSQL
│
▼
Unreal Game Servers

The backend acts as the **central authority for player data, abilities, and game infrastructure**.
