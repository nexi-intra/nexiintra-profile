package schemas

type Countries []struct {
	ID         int    `json:"ID"`
	Region     string `json:"Region"`
	RegionItem struct {
		LookupId    int    `json:"LookupId"`
		LookupValue string `json:"LookupValue"`
		TypeId      string `json:"TypeId"`
	} `json:"RegionItem"`
	Rolluppage struct {
		Description string `json:"Description"`
		TypeId      string `json:"TypeId"`
		Url         string `json:"Url"`
	} `json:"Rolluppage"`
	Title string `json:"Title"`
}
