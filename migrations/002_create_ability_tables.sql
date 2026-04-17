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