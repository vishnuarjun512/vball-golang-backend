CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create tables for the volleyball game backend
-- This includes tables for players, abilities, servers, bans, activity logs, regions, machines, game servers, and server players.

-- Players table to store player information and stats
CREATE TABLE players (
    player_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    steam_id TEXT UNIQUE NOT NULL,
    username TEXT NOT NULL,

    kash INTEGER DEFAULT 0,

    ban_status TEXT DEFAULT 'Active',

    matches_played INTEGER DEFAULT 0,
    wins INTEGER DEFAULT 0,
    losses INTEGER DEFAULT 0,

    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Main abilities table to store different main abilities that players can equip
CREATE TABLE main_abilities (
    id SERIAL PRIMARY KEY,

    name TEXT UNIQUE NOT NULL,
    description TEXT,

    type TEXT ,
    tier TEXT,

    duration FLOAT DEFAULT 0,
    cooldown FLOAT DEFAULT 0,

    spike_modifier FLOAT DEFAULT 1,
    jump_modifier FLOAT DEFAULT 1,
    set_modifier FLOAT DEFAULT 1,
    receive_modifier FLOAT DEFAULT 1,
    ball_force_multiplier FLOAT DEFAULT 1,

    created_at TIMESTAMP DEFAULT NOW()
);

-- Sub abilities table to store different sub abilities that players can equip
CREATE TABLE sub_abilities (
    id SERIAL PRIMARY KEY,

    name TEXT UNIQUE NOT NULL,
    description TEXT,

    tier TEXT,

    modifier_type TEXT,
    modifier_value FLOAT,

    created_at TIMESTAMP DEFAULT NOW()
);

-- Player abilities table to link players with their equipped main and sub abilities
CREATE TABLE player_abilities (

    player_id UUID PRIMARY KEY
        REFERENCES players(player_id)
        ON DELETE CASCADE,

    main_ability_id INTEGER
        REFERENCES main_abilities(id),

    sub_ability_slot1 INTEGER
        REFERENCES sub_abilities(id),

    sub_ability_slot2 INTEGER
        REFERENCES sub_abilities(id),

    sub_ability_slot3 INTEGER
        REFERENCES sub_abilities(id)
);

-- Servers table to track game servers and their status
CREATE TABLE servers (
    id TEXT PRIMARY KEY,
    region TEXT,
    max_players INTEGER,
    current_players INTEGER,
    game_mode TEXT,
    status TEXT,
    uptime TIMESTAMP
);

-- Bans table to track player bans and their details
CREATE TABLE bans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID REFERENCES players(player_id),
    reason TEXT,
    admin_id TEXT,
    ban_start TIMESTAMP,
    ban_end TIMESTAMP,
    type TEXT
);

-- Activity logs for tracking important events
CREATE TABLE activity_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type TEXT,
    message TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Regions
CREATE TABLE regions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Machines
CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    region_id INT REFERENCES regions(id) ON DELETE CASCADE,
    ip_address VARCHAR(100) UNIQUE NOT NULL,
    cpu_cores INT NOT NULL,
    ram_gb INT NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_machines_region
ON machines(region_id);

-- Game servers
CREATE TABLE game_servers (
    id SERIAL PRIMARY KEY,
    machine_id INT REFERENCES machines(id) ON DELETE CASCADE,
    port INT NOT NULL,
    max_players INT DEFAULT 20,
    current_players INT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'running',
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(machine_id, port)
);

CREATE INDEX IF NOT EXISTS idx_servers_status_players
ON game_servers(status, current_players);

-- Server players (NEW)
CREATE TABLE server_players (
    id SERIAL PRIMARY KEY,
    server_id INT REFERENCES game_servers(id) ON DELETE CASCADE,
    player_id VARCHAR(100) NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(server_id, player_id)
);

CREATE INDEX IF NOT EXISTS idx_server_players_player
ON server_players(player_id);