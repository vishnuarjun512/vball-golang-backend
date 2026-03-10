package models

type PlayerLoadout struct {
	PlayerID string `json:"playerId"`
	SteamID  string `json:"steamId"`
	Username string `json:"username"`
	Kash     int    `json:"kash"`

	MainAbilityID *int `json:"mainAbilityId"`

	SubAbilitySlot1 *int `json:"subAbilitySlot1"`
	SubAbilitySlot2 *int `json:"subAbilitySlot2"`
	SubAbilitySlot3 *int `json:"subAbilitySlot3"`
}
