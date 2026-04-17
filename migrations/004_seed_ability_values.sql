-- ==========================================================
-- 1. INDEPENDENT TABLES (Abilities & Regions)
-- ==========================================================

-- MAIN ABILITIES
INSERT INTO main_abilities 
(name, description, type, tier, duration, cooldown, spike_modifier, jump_modifier, set_modifier, receive_modifier, ball_force_multiplier)
VALUES
('StrongArm', 'Powerful spike ability.', 'Passive', 'Legendary', 0, 0, 1.35, 1.05, 1.0, 1.0, 1.2),
('ArcSetter', 'Higher precision arc sets.', 'Passive', 'Epic', 0, 0, 1.0, 1.0, 1.25, 1.0, 1.0),
('SoaringSpiker', 'Increases jump height.', 'Passive', 'Legendary', 0, 0, 1.2, 1.25, 1.0, 1.0, 1.1),
('SuperLibero', 'Exceptional receiving.', 'Passive', 'Epic', 0, 0, 1.0, 1.05, 1.0, 1.35, 1.0)
ON CONFLICT (name) DO NOTHING;

-- SUB ABILITIES
INSERT INTO sub_abilities (name, description, tier, modifier_type, modifier_value)
VALUES
('BigArms', 'Increases spike power.', 'Common', 'SpikePower', 1.10),
('FastApproach', 'Increases movement speed.', 'Common', 'MovementSpeed', 1.10),
('SkyStep', 'Improves jump height.', 'Rare', 'JumpHeight', 1.15),
('SnapSet', 'Improves set accuracy.', 'Rare', 'SetAccuracy', 1.20),
('SteelReceive', 'Improves receive accuracy.', 'Rare', 'ReceiveAccuracy', 1.25)
ON CONFLICT (name) DO NOTHING;

-- REGIONS
INSERT INTO regions (id, region_name, region_code)
VALUES
(1, 'North America - East', 'NA-EAST'),
(2, 'Europe - Central', 'EU-CENTRAL'),
(3, 'Asia - Southeast', 'ASIA-SE'),
(4, 'South America', 'SA-SOUTH')
ON CONFLICT (id) DO NOTHING;

-- ==========================================================
-- 2. INFRASTRUCTURE (Machines -> Game Servers)
-- ==========================================================

-- MACHINES (Depends on Regions)
INSERT INTO machines (id, region_id, machine_name, ip_address, cpu_cores, ram_gb, status, port_start, port_end, available_ports)
VALUES
(1, 2, 'EU - M1', '144.76.55.21', 8, 16, 'active', 7000, 7008, 8),
(2, 3, 'AS - M1', '103.21.44.12', 4, 8, 'stopped', 8000, 8004, 4),
(3, 1, 'NA - M1', '52.91.210.10', 8, 16, 'maintenance', 8200, 8208, 8),
(4, 2, 'EU - M2', '142.21.77.12', 16, 64, 'active', 6000, 6016, 16)
ON CONFLICT (id) DO NOTHING;

-- GAME SERVERS (Depends on Machines)
INSERT INTO game_servers (id, machine_id, port, max_players, current_players, status)
VALUES
(1, 1, 7000, 20, 5, 'running'),
(2, 1, 7001, 20, 10, 'running'),
(3, 1, 7002, 20, 18, 'running'),
(4, 2, 8000, 16, 0, 'stopped'),
(5, 3, 8200, 20, 3, 'running')
ON CONFLICT (id) DO NOTHING;

-- ==========================================================
-- 3. PLAYERS & SESSIONS (Depends on Game Servers)
-- ==========================================================

-- PLAYERS
INSERT INTO players (steam_id, username, kash, ban_status, matches_played, wins, losses, current_server)
VALUES 
('123', 'AceSpiker', 120, 'Active', 15, 10, 5, 1),
('456', 'SkyBlocker', 90, 'Active', 20, 12, 8, 3),
('789', 'SetMaster99', 50, 'Banned', 5, 1, 4, 2),
('101', 'NetRipper', 250, 'Active', 100, 75, 25, 1) 
ON CONFLICT (steam_id) DO NOTHING;

-- PLAYER ABILITIES (Connects Players to Ability IDs)
-- Note: Using subqueries to get player_id UUIDs safely
-- PLAYER ABILITIES
INSERT INTO player_abilities (player_id, main_ability_id, sub_ability_slot1, sub_ability_slot2, sub_ability_slot3)
SELECT player_id, 1, 1, 3, NULL::INTEGER FROM players WHERE steam_id = '123'
UNION ALL
SELECT player_id, 2, 2, 4, NULL FROM players WHERE steam_id = '456'
UNION ALL
SELECT player_id, 3, 1, 5, NULL FROM players WHERE steam_id = '789';

-- ==========================================================
-- 4. UTILITY TABLES (Ports)
-- ==========================================================
INSERT INTO ports (machine_id, port_number)
VALUES
(1, 7000), (1, 7001), (1, 7002), (1, 7003),
(2, 8000), (2, 8001), (2, 8002), (2, 8003),
(3, 8200), (3, 8201), (3, 8202), (3, 8203)
ON CONFLICT DO NOTHING;