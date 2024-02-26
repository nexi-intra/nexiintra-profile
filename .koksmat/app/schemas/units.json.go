package schemas

type Units []struct {
	ID          int         `json:"ID"`
	LookupValue interface{} `json:"LookupValue"`
	Rolluppage  struct {
		Description string `json:"Description"`
		TypeId      string `json:"TypeId"`
		Url         string `json:"Url"`
	} `json:"Rolluppage"`
	Site struct {
		Description string `json:"Description"`
		TypeId      string `json:"TypeId"`
		Url         string `json:"Url"`
	} `json:"Site"`
	SortOrder int    `json:"SortOrder"`
	Title     string `json:"Title"`
	UnitType  string `json:"UnitType"`
}
