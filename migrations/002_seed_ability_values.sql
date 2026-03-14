-- MAIN ABILITIES

INSERT INTO main_abilities
(name, description, type, tier, duration, cooldown,
spike_modifier, jump_modifier, set_modifier, receive_modifier, ball_force_multiplier)

VALUES

(
'StrongArm',
'Powerful spike ability that increases spike force.',
'Passive',
'Legendary',
0,
0,
1.35,
1.05,
1.0,
1.0,
1.2
),

(
'ArcSetter',
'Allows higher precision arc sets.',
'Passive',
'Epic',
0,
0,
1.0,
1.0,
1.25,
1.0,
1.0
),

(
'SoaringSpiker',
'Greatly increases jump height for spikes.',
'Passive',
'Legendary',
0,
0,
1.2,
1.25,
1.0,
1.0,
1.1
),

(
'SuperLibero',
'Exceptional receiving ability.',
'Passive',
'Epic',
0,
0,
1.0,
1.05,
1.0,
1.35,
1.0
)

ON CONFLICT (name) DO NOTHING;

-- SUB ABILITIES

INSERT INTO sub_abilities
(name, description, tier, modifier_type, modifier_value)

VALUES

(
'BigArms',
'Increases spike power slightly.',
'Common',
'SpikePower',
1.10
),

(
'FastApproach',
'Increases player movement speed.',
'Common',
'MovementSpeed',
1.10
),

(
'SkyStep',
'Improves jump height.',
'Rare',
'JumpHeight',
1.15
),

(
'SnapSet',
'Improves set accuracy.',
'Rare',
'SetAccuracy',
1.20
),

(
'SteelReceive',
'Improves receive accuracy.',
'Rare',
'ReceiveAccuracy',
1.25
)

ON CONFLICT (name) DO NOTHING;


-- SERVERS

INSERT INTO servers
(id, region, max_players, current_players, game_mode, status)
VALUES

('S-001','mumbai',30,18,'Ranked','Online'),

('S-002','singapore',30,12,'Ranked','Online'),

('S-003','na-east',30,0,'Casual','Offline');

-- PLAYERS

INSERT INTO players
(steam_id, username, kash)
VALUES

('123','AceSpiker',120),
('456','SkyBlocker',90),
('789','SetMaster99',50)

ON CONFLICT (steam_id) DO NOTHING;

-- PLAYER ABILITIES

INSERT INTO player_abilities
(player_id, main_ability_id, sub_ability_slot1, sub_ability_slot2, sub_ability_slot3)

SELECT player_id, 1, 1, 3, NULL::INTEGER
FROM players WHERE steam_id = '123'

UNION ALL

SELECT player_id, 2, 2, 4, NULL::INTEGER
FROM players WHERE steam_id = '456'

UNION ALL

SELECT player_id, 3, 1, 5, NULL::INTEGER
FROM players WHERE steam_id = '789';


-- Regions
INSERT INTO regions (name)
VALUES ('eu'), ('asia'), ('us')
ON CONFLICT DO NOTHING;

-- VPS Machines:
INSERT INTO machines (region_id, ip_address, cpu_cores, ram_gb)
VALUES
(1,'144.76.55.21',8,16),
(2,'103.21.44.12',8,16),
(3,'52.91.210.10',8,16)
ON CONFLICT DO NOTHING;

-- Game servers
INSERT INTO game_servers (machine_id, port, max_players)
VALUES
(1,7777,20),
(1,7778,20),
(2,7777,20),
(2,7778,20),
(3,7777,20)
ON CONFLICT DO NOTHING;