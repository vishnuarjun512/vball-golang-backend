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
    ball_force_modifier FLOAT DEFAULT 1,

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