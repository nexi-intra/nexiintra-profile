package schemas

type Categories []struct {
	ID        int    `json:"ID"`
	SortOrder int    `json:"SortOrder"`
	Title     string `json:"Title"`
}
