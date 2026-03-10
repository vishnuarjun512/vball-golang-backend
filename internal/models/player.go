package models

import "time"

type Player struct {
	PlayerID     string   `json:"player_id"`
	Username     string   `json:"username"`
	MainAbility  string   `json:"main_ability"`
	SubAbilities []string `json:"sub_abilities"`
}

type PlayerAdmin struct {
	PlayerID string `json:"playerId"`
	SteamID  string `json:"steamId"`
	Username string `json:"username"`

	Kash int `json:"kash"`

	MainAbilityID *int  `json:"mainAbilityId"`
	SubAbilityIDs []int `json:"subAbilityIds"`

	BanStatus string `json:"banStatus"`

	MatchesPlayed int `json:"matchesPlayed"`
	Wins          int `json:"wins"`
	Losses        int `json:"losses"`

	LastLogin    *time.Time `json:"lastLogin"`
	RegisteredAt time.Time  `json:"registeredAt"`
}
