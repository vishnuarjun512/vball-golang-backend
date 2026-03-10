# API Routes

This document lists the available API routes.

---

## Authentication

```
POST /auth/steam-login
```

Registers a player automatically if they do not exist.

---

## Player Routes

```
GET /admin/players
GET /admin/players/:steamid
```

Returns player data including ability IDs.

---

## Ability Routes

```
GET /admin/abilities
GET /admin/main-abilities
GET /admin/sub-abilities
```

Returns ability definitions.

---

## Game Server Routes

These endpoints are used by Unreal Engine servers.

```
GET /game/player/:steamid/loadout
GET /game/abilities
```

Game servers receive **only ability IDs** to keep network payloads small.
