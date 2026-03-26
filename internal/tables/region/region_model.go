package region

type RegionReq struct {
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
}

type Region struct {
	ID         int    `db:"id" json:"id"`
	RegionName string `db:"region_name" json:"region_name"`
	RegionCode string `db:"region_code" json:"region_code"`
}
