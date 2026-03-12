package models

type CreateAbilityRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tier string `json:"tier"`

	Duration int `json:"duration"`
	Cooldown int `json:"cooldown"`

	Description string `json:"description"`

	SpikeModifier       float64 `json:"spike_modifier"`
	JumpModifier        float64 `json:"jump_modifier"`
	SetModifier         float64 `json:"set_modifier"`
	ReceiveModifier     float64 `json:"receive_modifier"`
	BallForceMultiplier float64 `json:"ball_force_multiplier"`
}

type MainAbility struct {
	ID int `db:"id" json:"id"`

	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`

	Type string `db:"type" json:"type"`
	Tier string `db:"tier" json:"tier"`

	Duration float64 `db:"duration" json:"duration"`
	Cooldown float64 `db:"cooldown" json:"cooldown"`

	SpikeModifier       float64 `db:"spike_modifier" json:"spike_modifier"`
	JumpModifier        float64 `db:"jump_modifier" json:"jump_modifier"`
	SetModifier         float64 `db:"set_modifier" json:"set_modifier"`
	ReceiveModifier     float64 `db:"receive_modifier" json:"receive_modifier"`
	BallForceMultiplier float64 `db:"ball_force_multiplier" json:"ball_force_multiplier"`
}

type SubAbility struct {
	ID int `db:"id" json:"id"`

	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`

	Tier string `db:"tier" json:"tier"`

	ModifierType  string  `db:"modifier_type" json:"modifier_type"`
	ModifierValue float64 `db:"modifier_value" json:"modifier_value"`
}
