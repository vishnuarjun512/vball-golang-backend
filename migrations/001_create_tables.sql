CREATE EXTENSION IF NOT EXISTS "pgcrypto";

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

CREATE TABLE sub_abilities (
    id SERIAL PRIMARY KEY,

    name TEXT UNIQUE NOT NULL,
    description TEXT,

    tier TEXT,

    modifier_type TEXT,
    modifier_value FLOAT,

    created_at TIMESTAMP DEFAULT NOW()
);

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

CREATE TABLE servers (
    id TEXT PRIMARY KEY,
    region TEXT,
    max_players INTEGER,
    current_players INTEGER,
    game_mode TEXT,
    status TEXT,
    uptime TIMESTAMP
);

CREATE TABLE bans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID REFERENCES players(player_id),
    reason TEXT,
    admin_id TEXT,
    ban_start TIMESTAMP,
    ban_end TIMESTAMP,
    type TEXT
);

CREATE TABLE activity_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type TEXT,
    message TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Regions
CREATE TABLE regions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Machines (VPS / physical servers)
CREATE TABLE machines (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    region_id UUID REFERENCES regions(id) ON DELETE CASCADE,
    ip_address VARCHAR(100) UNIQUE NOT NULL,
    cpu_cores INT NOT NULL,
    ram_gb INT NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW()
);

-- Game server instances
CREATE TABLE game_servers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    machine_id UUID REFERENCES machines(id) ON DELETE CASCADE,
    port INT NOT NULL,
    max_players INT DEFAULT 20,
    current_players INT DEFAULT 0,
    status VARCHAR(20) DEFAULT 'waiting',
    match_id UUID REFERENCES matches(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(machine_id, port)
);

-- Matches
CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    server_id UUID REFERENCES game_servers(id),
    status VARCHAR(20) DEFAULT 'waiting',
    created_at TIMESTAMP DEFAULT NOW()
);

-- Match players
CREATE TABLE match_players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    match_id UUID REFERENCES matches(id) ON DELETE CASCADE,
    player_id VARCHAR(100) NOT NULL,
    team INT,
    joined_at TIMESTAMP DEFAULT NOW()
);