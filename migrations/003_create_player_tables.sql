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

    -- Added ON DELETE SET NULL
    current_server INT REFERENCES game_servers(id) ON DELETE SET NULL,

    last_login TIMESTAMP,
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

