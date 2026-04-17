CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create tables for the volleyball game backend
-- This includes tables for players, abilities, servers, bans, activity logs, regions, machines, game servers, and server players.

-- Regions
CREATE TABLE regions (
    id SERIAL PRIMARY KEY,
    region_name VARCHAR(50) UNIQUE NOT NULL,
    region_code VARCHAR(50)
);

-- Machines
CREATE TABLE machines (
    id SERIAL PRIMARY KEY,
    machine_name VARCHAR(20),
    region_id INT REFERENCES regions(id) ON DELETE CASCADE,
    ip_address VARCHAR(100) UNIQUE NOT NULL,
    cpu_cores INT NOT NULL,
    ram_gb INT NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    port_start INT NOT NULL,
    port_end INT NOT NULL,
    available_ports INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_machines_region
ON machines(region_id);

-- PORTS
CREATE TABLE ports (
    id SERIAL PRIMARY KEY,
    machine_id INT REFERENCES machines(id) ON DELETE CASCADE,
    port_number INT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_ports_machine_id
ON ports(machine_id);

-- Game servers
CREATE TABLE game_servers (
    id SERIAL PRIMARY KEY,
    machine_id INT REFERENCES machines(id) ON DELETE CASCADE,
    port INT NOT NULL,
    max_players INT DEFAULT 20,
    current_players INT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'running',
    uptime TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(machine_id, port)
);

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

