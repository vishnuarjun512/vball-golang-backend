# Volleyball Game Backend

Backend service for a **multiplayer volleyball game** built with Go.

This backend powers:

- Player accounts and Steam authentication
- Ability + sub-ability systems
- Dedicated game server management
- Matchmaking and player-server assignment
- Admin dashboard APIs

The system is designed to work with **Unreal Engine multiplayer servers** while providing a **web-based admin panel** for managing players, abilities, and servers.

---

## Tech Stack

| Concern          | Technology       |
| ---------------- | ---------------- |
| Language         | Go (Golang)      |
| Framework        | Gin              |
| Database         | PostgreSQL       |
| Driver           | pgx              |
| Containerization | Docker           |
| Dev Tooling      | Air (hot reload) |
| Game Engine      | Unreal Engine    |
| Auth             | Steam            |

---

## Features

- Steam-based player authentication and auto-registration
- Ability + sub-ability system with tiers and modifiers
- Player loadout management
- Matchmaking — assigns players to available regional servers
- Admin player and ability management
- Game server loadout APIs (ability IDs only, lightweight payloads)
- Infrastructure tracking — regions, machines, game server instances
- Moderation — bans and activity logs
- Dockerized development environment
- PostgreSQL migrations + seed data

---

## Installation

Clone the repository.

```bash
git clone https://github.com/yourusername/vball-go-backend.git
cd vball-go-backend
```

Start the backend.

```bash
docker compose up --build
```

The API will be available at: `http://localhost:8080`

---

## Development

Run the backend locally with hot reload.

```bash
docker compose up
```

Reset the database.

```bash
docker compose down -v
docker compose up --build
```

---

## Project Structure

```
cmd/
    api/            → application entrypoint
internal/
    database/       → database connection
    handlers/       → HTTP handlers
    services/       → business logic
    repositories/   → database queries
    models/         → data models
    routes/         → API route definitions
migrations/         → database schema + seed files
docker/             → container configuration
docs/               → project documentation
    architecture.md
    api-routes.md
    database.md
```

---

## API Overview

Base URL: `http://localhost:8080`

| Method | Route                           | Description                        |
| ------ | ------------------------------- | ---------------------------------- |
| POST   | `/auth/steam-login`             | Authenticate and register player   |
| GET    | `/admin/players`                | List all players                   |
| GET    | `/admin/players/:steamid`       | Get player + loadout               |
| GET    | `/admin/abilities`              | List all abilities                 |
| GET    | `/abilities/main`               | List main abilities                |
| POST   | `/abilities/main`               | Create main ability                |
| PATCH  | `/abilities/main/:id`           | Update main ability                |
| DELETE | `/abilities/main/:id`           | Delete main ability                |
| GET    | `/game/player/:steamid/loadout` | Game server: fetch player loadout  |
| GET    | `/game/abilities`               | Game server: fetch ability list    |
| POST   | `/game/matchmaking/join`        | Player joins matchmaking queue     |
| POST   | `/game/matchmaking/leave`       | Player leaves or disconnects       |
| POST   | `/game/matchmaking/sync`        | Server syncs connected player list |

See [docs/api-routes.md](docs/api-routes.md) for full request/response details.

---

## Documentation

All documentation lives in the `/docs` directory.

| Document                                     | Description                                                                 |
| -------------------------------------------- | --------------------------------------------------------------------------- |
| [docs/architecture.md](docs/architecture.md) | System design, request flows, infrastructure layout, and technology choices |
| [docs/api-routes.md](docs/api-routes.md)     | All API endpoints with example requests and responses                       |
| [docs/database.md](docs/database.md)         | Database schema, table definitions, indexes, and psql reference             |

---

## Roadmap

Planned features include:

- Matchmaking system (in progress)
- Server orchestration
- Ranked ladder
- Ability reroll mechanics
- Match history
- Player statistics

---

## Contributing

Contributions are welcome.

If you'd like to improve the backend:

- Fork the repository
- Create a feature branch
- Submit a pull request

---

## License

MIT License
