# API Routes

This document describes the available backend API endpoints.

The backend exposes APIs for:

- Player management
- Ability management
- Matchmaking
- Admin operations

Base URL:

```
http://localhost:8080
```

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

### Get Player Loadout

```
GET /admin/players/:steamid
```

Returns a player's loadout including abilities.

Example response:

```json
{
  "steamId": "steam_123",
  "username": "AceSpiker",
  "mainAbility": "StrongArm",
  "subAbilities": ["QuickStep", "IronDefense"]
}
```

---

## Ability Routes

```
GET /admin/abilities
GET /admin/main-abilities
GET /admin/sub-abilities
```

Returns ability definitions.

### Get All Main Abilities

```
GET /abilities/main
```

Returns all main abilities available in the game.

### Create Main Ability

```
POST /abilities/main
```

Example request:

```json
{
  "name": "StrongArm",
  "type": "Passive",
  "tier": "Legendary",
  "spike_modifier": 1.5
}
```

### Update Ability

```
PATCH /abilities/main/:id
```

### Delete Ability

```
DELETE /abilities/main/:id
```

---

## Game Server Routes

These endpoints are used by Unreal Engine servers.

```
GET /game/player/:steamid/loadout
GET /game/abilities
```

Game servers receive **only ability IDs** to keep network payloads small.

---

## Matchmaking Routes

### Join Matchmaking

```
POST /game/matchmaking/join
```

Used by the Unreal client when a player presses Play.

Example request:

```json
{
  "playerId": "steam_999",
  "region": "asia"
}
```

Backend actions:

1. Find available server
2. Assign player
3. Increment server player count

Example response:

```json
{
  "serverIp": "103.21.44.12",
  "port": 7777
}
```

### Leave Server

```
POST /game/matchmaking/leave
```

Used when a player leaves or disconnects.

Example request:

```json
{
  "playerId": "steam_999"
}
```

Backend actions:

- Remove player from `server_players`
- Decrement `current_players` count on the server

### Sync Server Players

```
POST /game/matchmaking/sync
```

Used by game servers to synchronize their connected player list.

Example request:

```json
{
  "serverId": 1,
  "players": ["steam_1", "steam_2", "steam_3"]
}
```

Backend updates:

- `server_players` table
- `current_players` count
