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
