package unitmodel

type Unit struct {
	Project     string `json:"project"`
	UnitID      string `json:"unitid"`
	Duration    int    `json:"duration"`
	Value       string `json:"value"`
}

type UnitNoTime struct {
	Project     string `json:"project"`
	UnitID      string `json:"unitid"`
	Value       string `json:"value"`
}

type UnitReq struct {
	Project     string `json:"project"`
	UnitID      string `json:"unitid"`
}
