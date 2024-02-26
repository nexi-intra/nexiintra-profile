package schemas

type Profiledata struct {
	Categories []struct {
		ID        int    `json:"ID"`
		SortOrder int    `json:"SortOrder"`
		Title     string `json:"Title"`
	} `json:"categories"`
	Channels []struct {
		ID           int  `json:"ID"`
		Mandatory    bool `json:"Mandatory"`
		NewsCategory struct {
			LookupId    int    `json:"LookupId"`
			LookupValue string `json:"LookupValue"`
			TypeId      string `json:"TypeId"`
		} `json:"NewsCategory"`
		RelevantCountires []interface{} `json:"RelevantCountires"`
		RelevantUnits     []interface{} `json:"RelevantUnits"`
		GroupID           string        `json:"GroupId"`
		Title             string        `json:"Title"`
		RegionID          int           `json:"RegionId"`
		NewsCategoryID    int           `json:"NewsCategoryId"`
	} `json:"channels"`
	Countries []struct {
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
	} `json:"countries"`
	Units []struct {
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
	} `json:"units"`
}
