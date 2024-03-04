package schemas

type Washedprofiledata struct {
	_Id struct {
		_Oid string `json:"$oid"`
	} `json:"_id"`
	Categories []struct {
		CategoryId   int    `json:"categoryId"`
		CategoryName string `json:"categoryName"`
		SortOrder    int    `json:"sortOrder"`
	} `json:"categories"`
	Channels []struct {
		GroupId           string        `json:"GroupId"`
		Mandatory         bool          `json:"Mandatory"`
		NewsCategoryId    int           `json:"NewsCategoryId"`
		RegionId          int           `json:"RegionId"`
		RelevantCountires []interface{} `json:"RelevantCountires"`
		RelevantUnits     []interface{} `json:"RelevantUnits"`
		ChannelCode       string        `json:"channelCode"`
		ChannelName       string        `json:"channelName"`
		ChannelType       string        `json:"channelType"`
		Id                string        `json:"id"`
		SortOrder         string        `json:"sortOrder"`
	} `json:"channels"`
	Countries []struct {
		CountryCode string `json:"countryCode"`
		CountryName string `json:"countryName"`
	} `json:"countries"`
	Date struct {
		_Date string `json:"$date"`
	} `json:"date"`
	Key   string `json:"key"`
	Units []struct {
		SortOrder int    `json:"sortOrder"`
		UnitCode  string `json:"unitCode"`
		UnitName  string `json:"unitName"`
		UnitType  string `json:"unitType"`
	} `json:"units"`
}
