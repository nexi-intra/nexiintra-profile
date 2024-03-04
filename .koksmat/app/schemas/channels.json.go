package schemas

type Channels []struct {
	GroupId      string `json:"GroupId"`
	ID           int    `json:"ID"`
	Mandatory    bool   `json:"Mandatory"`
	NewsCategory struct {
		LookupId    int    `json:"LookupId"`
		LookupValue string `json:"LookupValue"`
		TypeId      string `json:"TypeId"`
	} `json:"NewsCategory"`
	NewsCategoryID    int           `json:"NewsCategoryID"`
	RegionID          interface{}   `json:"RegionID"`
	RelevantCountires []interface{} `json:"RelevantCountires"`
	RelevantUnits     []interface{} `json:"RelevantUnits"`
	Title             string        `json:"Title"`
}
