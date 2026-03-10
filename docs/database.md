# PostgreSQL Commands

Connect to your container:

```
docker exec -it vball-postgres psql -U postgres -d volleyball
```

Then inside:

```
\dt
\d players
SELECT * FROM players;
```

### Navigation Commands

- List all databases

```
\l
```

or

```
\list
```

### Connect to another database

```
\c database_name
```

Example:

```
\c volleyball
```

### Show current database

```
SELECT current_database();
```

## Table Inspection

### Show tables

```
\dt
```

### Show table structure

```
\d table_name
```

Example:

```
\d players
```

Shows:

- columns
- types
- indexes
- constraints

Show indexes:

- \di

Show sequences:

- \ds

Show views:

- \dv

3. Query Data
   Show all rows

```
SELECT * FROM players;
```

Limit results

```
SELECT * FROM players LIMIT 10
```

Filter rows

```
SELECT * FROM players WHERE steam_id = 123;
```

Count rows

```
SELECT COUNT(*) FROM players;
```

### Insert / Update / Delete

##### Insert player

```
INSERT INTO players (steam_id, username)
VALUES (123456, 'Vishnu');
```

##### Update player

```
UPDATE players
SET username = 'NewName'
WHERE steam_id = 123456;
```

##### Delete player

```
DELETE FROM players
WHERE steam_id = 123456;
```

### Table Management

##### Create table

```
CREATE TABLE players (
    steam_id BIGINT PRIMARY KEY,
    username TEXT
);
```

##### Delete table

```
DROP TABLE players;
Empty table (keep structure)
TRUNCATE TABLE players;
```

### Search in tables

Very useful command:

```
\d
```

Shows all relations.

#### Show schemas

```
\dn
```

#### Show roles (users)

```
\du
```

#### Command history

```
Press:
↑ arrow key
```

to see previous commands.

#### Help commands

Show psql help:

```
\?
```

Show SQL help:

```
\h
```

Example:

```
\h SELECT
```

### Formatting output

Turn expanded output on (very useful):

```
\x
```

Now large rows display vertically.

### Exit psql

```
\q
```

# Database Schema

## players

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

## player_abilities

Links players to abilities.

| Column            | Type    |
| ----------------- | ------- |
| player_id         | UUID    |
| main_ability_id   | INTEGER |
| sub_ability_slot1 | INTEGER |
| sub_ability_slot2 | INTEGER |
| sub_ability_slot3 | INTEGER |

---

## main_abilities

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

## sub_abilities

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

## servers

Tracks active game servers.

| Column          | Type    |
| --------------- | ------- |
| id              | TEXT    |
| region          | TEXT    |
| max_players     | INTEGER |
| current_players | INTEGER |
| game_mode       | TEXT    |
| status          | TEXT    |

---

## bans

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

## activity_logs

Tracks system events.

| Column     | Type      |
| ---------- | --------- |
| id         | UUID      |
| type       | TEXT      |
| message    | TEXT      |
| created_at | TIMESTAMP |
