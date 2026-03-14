# Database Schema

This document describes the PostgreSQL database structure used by the Volleyball Game Backend.

The database stores:

- Infrastructure information
- Player assignments
- Abilities
- Player loadouts

---

## PostgreSQL Commands

Connect to your container:

```bash
docker exec -it vball-postgres psql -U postgres -d volleyball
```

Then inside:

```sql
\dt
\d players
SELECT * FROM players;
```

### Navigation

List all databases:

```
\l
```

Connect to another database:

```
\c database_name
```

Show current database:

```sql
SELECT current_database();
```

### Table Inspection

Show all tables:

```
\dt
```

Show table structure:

```
\d table_name
```

Shows columns, types, indexes, and constraints.

Other useful inspection commands:

- `\di` — show indexes
- `\ds` — show sequences
- `\dv` — show views
- `\dn` — show schemas
- `\du` — show roles (users)
- `\d` — show all relations

### Querying Data

Show all rows:

```sql
SELECT * FROM players;
```

Limit results:

```sql
SELECT * FROM players LIMIT 10;
```

Filter rows:

```sql
SELECT * FROM players WHERE steam_id = 123;
```

Count rows:

```sql
SELECT COUNT(*) FROM players;
```

### Insert / Update / Delete

Insert a player:

```sql
INSERT INTO players (steam_id, username)
VALUES (123456, 'Vishnu');
```

Update a player:

```sql
UPDATE players
SET username = 'NewName'
WHERE steam_id = 123456;
```

Delete a player:

```sql
DELETE FROM players
WHERE steam_id = 123456;
```

### Table Management

Create a table:

```sql
CREATE TABLE players (
    steam_id BIGINT PRIMARY KEY,
    username TEXT
);
```

Delete a table:

```sql
DROP TABLE players;
```

Empty a table (keep structure):

```sql
TRUNCATE TABLE players;
```

### Formatting & Help

Turn on expanded output (useful for wide rows):

```
\x
```

Show psql help:

```
\?
```

Show SQL help:

```
\h SELECT
```

View command history — press the `↑` arrow key.

Exit psql:

```
\q
```

---

## Infrastructure Tables

### Regions

Stores available server regions.

| Column | Type               |
| ------ | ------------------ |
| id     | SERIAL PRIMARY KEY |
| name   | VARCHAR(50) UNIQUE |

Example rows: `1 eu`, `2 asia`, `3 us`

---

### Machines

Represents VPS or physical servers.

| Column     | Type                       |
| ---------- | -------------------------- |
| id         | SERIAL PRIMARY KEY         |
| region_id  | INT REFERENCES regions(id) |
| ip_address | VARCHAR(100)               |
| cpu_cores  | INT                        |
| ram_gb     | INT                        |
| status     | VARCHAR(20)                |
| created_at | TIMESTAMP                  |

Example rows: `1 | region 1 | 144.76.55.21`, `2 | region 2 | 103.21.44.12`

---

### Game Servers

Represents Unreal Engine server instances running on machines.

| Column          | Type                        |
| --------------- | --------------------------- |
| id              | SERIAL PRIMARY KEY          |
| machine_id      | INT REFERENCES machines(id) |
| port            | INT                         |
| max_players     | INT                         |
| current_players | INT                         |
| status          | VARCHAR(20)                 |
| created_at      | TIMESTAMP                   |

Unique constraint: `UNIQUE(machine_id, port)`

Each server is identified by `machine_ip + port`, for example: `103.21.44.12:7777`

---

### Server Players

Tracks players connected to game servers.

| Column    | Type                            |
| --------- | ------------------------------- |
| id        | SERIAL PRIMARY KEY              |
| server_id | INT REFERENCES game_servers(id) |
| player_id | VARCHAR(100)                    |
| joined_at | TIMESTAMP                       |

---

## Player Tables

### players

Stores player account data.

| Column         | Type      |
| -------------- | --------- |
| player_id      | UUID      |
| steam_id       | TEXT      |
| username       | TEXT      |
| kash           | INTEGER   |
| ban_status     | TEXT      |
| matches_played | INTEGER   |
| wins           | INTEGER   |
| losses         | INTEGER   |
| last_login     | TIMESTAMP |
| created_at     | TIMESTAMP |

---

### player_abilities

Links players to their ability loadout.

| Column            | Type    |
| ----------------- | ------- |
| player_id         | UUID    |
| main_ability_id   | INTEGER |
| sub_ability_slot1 | INTEGER |
| sub_ability_slot2 | INTEGER |
| sub_ability_slot3 | INTEGER |

---

## Ability Tables

### main_abilities

Defines primary abilities.

| Column              | Type   |
| ------------------- | ------ |
| id                  | SERIAL |
| name                | TEXT   |
| description         | TEXT   |
| ability_type        | TEXT   |
| tier                | TEXT   |
| duration            | FLOAT  |
| cooldown            | FLOAT  |
| spike_modifier      | FLOAT  |
| jump_modifier       | FLOAT  |
| set_modifier        | FLOAT  |
| receive_modifier    | FLOAT  |
| ball_force_modifier | FLOAT  |

---

### sub_abilities

Defines modifier abilities.

| Column         | Type   |
| -------------- | ------ |
| id             | SERIAL |
| name           | TEXT   |
| description    | TEXT   |
| tier           | TEXT   |
| modifier_type  | TEXT   |
| modifier_value | FLOAT  |

---

## Moderation Tables

### bans

Stores moderation actions.

| Column    | Type      |
| --------- | --------- |
| id        | UUID      |
| player_id | UUID      |
| reason    | TEXT      |
| admin_id  | TEXT      |
| ban_start | TIMESTAMP |
| ban_end   | TIMESTAMP |
| type      | TEXT      |

---

### activity_logs

Tracks system events.

| Column     | Type      |
| ---------- | --------- |
| id         | UUID      |
| type       | TEXT      |
| message    | TEXT      |
| created_at | TIMESTAMP |

---

## Database Indexes

Indexes are added to optimize matchmaking and lookup queries.

| Index Name                   | Table          | Columns                   |
| ---------------------------- | -------------- | ------------------------- |
| `idx_servers_status_players` | game_servers   | (status, current_players) |
| `idx_machines_region`        | machines       | (region_id)               |
| `idx_server_players_player`  | server_players | (player_id)               |

These indexes allow fast lookups when:

- Finding available servers
- Locating players across servers
- Filtering servers by region
