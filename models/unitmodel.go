package unitmodel

type Unit struct {
	Project     string `json:"title"`
	UnitID      string `json:"unitid"`
	Duration    int    `json:"duration"`
	Value       string `json:"value"`
}
