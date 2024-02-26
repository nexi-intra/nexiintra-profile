package schemas

type Channels []struct {
	ID           int  `json:"ID"`
	Mandatory    bool `json:"Mandatory"`
	NewsCategory struct {
		LookupId    int    `json:"LookupId"`
		LookupValue string `json:"LookupValue"`
		TypeId      string `json:"TypeId"`
	} `json:"NewsCategory"`
	RelevantCountires []interface{} `json:"RelevantCountires"`
	RelevantUnits     []interface{} `json:"RelevantUnits"`
	Title             string        `json:"Title"`
}
