package subAbility

type SubAbility struct {
	ID int `db:"id" json:"id"`

	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`

	Tier string `db:"tier" json:"tier"`

	ModifierType  string  `db:"modifier_type" json:"modifier_type"`
	ModifierValue float64 `db:"modifier_value" json:"modifier_value"`
}
