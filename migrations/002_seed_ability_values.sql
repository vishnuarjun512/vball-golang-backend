-- MAIN ABILITIES

INSERT INTO main_abilities
(name, description, ability_type, tier, duration, cooldown,
spike_modifier, jump_modifier, set_modifier, receive_modifier, ball_force_modifier)

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