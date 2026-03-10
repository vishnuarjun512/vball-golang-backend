# Backend Architecture

The Volleyball Backend uses a layered architecture to keep the codebase modular and scalable.

## Request Flow

Client Request
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

## Layer Responsibilities

### Routes

Defines API endpoints.

### Handlers

Processes HTTP requests and responses.

### Services

Contains business logic.

### Repositories

Handles database queries.

### Database

Stores persistent game data.

---

## System Overview

Admin Dashboard
↓
Go Backend API
↓
PostgreSQL
↓
Unreal Dedicated Game Servers

The backend acts as the **central authority for player data and abilities**.

Game servers query the backend for player loadouts when players join matches.

---

## Ability System

Abilities are defined separately from players.

Tables:

- `main_abilities`
- `sub_abilities`

Players reference abilities through:

player_abilities

This allows:

- ability balancing without updating players
- small API responses
- reusable ability definitions

---

## Player Registration Flow

Player launches game
↓
Steam authentication
↓
POST /auth/steam-login
↓
Backend checks if player exists
↓
Creates player if necessary
↓
Returns player data

This allows **automatic account creation**.
