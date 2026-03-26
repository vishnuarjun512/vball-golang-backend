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
INSERT INTO regions (id, region_name, region_code)
VALUES
(1,'North America - East','NA-EAST'),
(2,'Europe - Central','EU-CENTRAL'),
(3,'Asia - Southeast','ASIA-SE'),
(4,'South America','SA-SOUTH')
ON CONFLICT DO NOTHING;


-- VPS Machines:
INSERT INTO machines 
(region_id, machine_name, ip_address, cpu_cores, ram_gb, status, port_start, port_end, available_ports)
VALUES
(2, 'EU - M1', '144.76.55.21', 8, 16, 'active', 7000, 7008, 8),
(3, 'AS - M1', '103.21.44.12', 4, 8, 'stopped', 8000, 8004, 4),
(1, 'NA - M1', '52.91.210.10', 8, 16, 'maintenance', 8200, 8208, 8),
(2, 'EU - M2', '142.21.77.12', 16, 64, 'active', 6000, 6016, 16)
ON CONFLICT DO NOTHING;

-- Ports
INSERT INTO ports (machine_id, port_number)
VALUES

-- Machine 1
(1,7000),(1,7001),(1,7002),(1,7003),
(1,7004),(1,7005),(1,7006),(1,7007),

-- Machine 2
(2,8000),(2,8001),(2,8002),(2,8003),

-- Machine 3
(3,8200),(3,8201),(3,8202),(3,8203),
(3,8204),(3,8205),(3,8206),(3,8207),

-- Machine 4
(4,6000),(4,6001),(4,6002),(4,6003),
(4,6004),(4,6005),(4,6006),(4,6007),
(4,6008),(4,6009),(4,6010),(4,6011),
(4,6012),(4,6013),(4,6014),(4,6015)
ON CONFLICT DO NOTHING;

-- Game Servers
INSERT INTO game_servers
(machine_id, port, max_players, current_players, status)
VALUES

(1,7000,20,5,'running'),
(1,7001,20,10,'running'),
(1,7002,20,18,'running'),

(2,8000,16,0,'stopped'),

(3,8200,20,3,'running'),
(3,8201,20,7,'running'),

(4,6000,20,12,'running'),
(4,6001,20,18,'running'),
(4,6002,20,20,'full'),
(4,6003,20,8,'running')
ON CONFLICT DO NOTHING;

-- Players
INSERT INTO server_players (server_id, player_id)
VALUES
(1,'player_1001'),
(1,'player_1002'),
(1,'player_1003'),

(2,'player_2001'),
(2,'player_2002'),

(4,'player_3001'),
(4,'player_3002'),
(4,'player_3003'),

(5,'player_4001'),
(5,'player_4002'),
(5,'player_4003'),
(5,'player_4004')
ON CONFLICT DO NOTHING;