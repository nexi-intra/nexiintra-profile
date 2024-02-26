package execution

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/365admin/nexiintra-profile/schemas"
)

type Country struct {
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
}

type Category struct {
	CategoryName string `json:"categoryName"`
	CategoryID   int    `json:"categoryId"`
	SortOrder    int    `json:"sortOrder"`
}

type Unit struct {
	UnitName  string `json:"unitName"`
	UnitCode  string `json:"unitCode"`
	UnitType  string `json:"unitType"`
	SortOrder int    `json:"sortOrder"`
}

type Channel struct {
	ID                string `json:"id"`
	SortOrder         string `json:"sortOrder"`
	ChannelName       string `json:"channelName"`
	ChannelType       string `json:"channelType"`
	ChannelCode       string `json:"channelCode"`
	GroupID           string `json:"GroupId"`
	RelevantUnits     []any  `json:"RelevantUnits"`
	Mandatory         bool   `json:"Mandatory"`
	RelevantCountires []any  `json:"RelevantCountires"`
	RegionID          int    `json:"RegionId"`
	NewsCategoryID    int    `json:"NewsCategoryId"`
}
type ProfileData struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Date struct {
		Date time.Time `json:"$date"`
	} `json:"date"`
	Key        string     `json:"key"`
	Countries  []Country  `json:"countries"`
	Categories []Category `json:"categories"`
	Units      []Unit     `json:"units"`
	Channels   []Channel  `json:"channels"`
}

func WashProfileData(profileData *schemas.Profiledata) (ProfileData, error) {
	// Wash the data
	var washedData ProfileData
	for _, country := range profileData.Countries {
		washedData.Countries = append(washedData.Countries, Country{
			CountryName: country.Title,
			CountryCode: country.Title,
		})
	}
	for _, category := range profileData.Categories {
		washedData.Categories = append(washedData.Categories, Category{
			CategoryName: category.Title,
			CategoryID:   category.ID,
			SortOrder:    category.SortOrder,
		})
	}

	for _, unit := range profileData.Units {
		washedData.Units = append(washedData.Units, Unit{
			UnitName:  unit.Title,
			UnitCode:  unit.Title,
			UnitType:  unit.UnitType,
			SortOrder: unit.SortOrder,
		})
	}

	for _, channel := range profileData.Channels {
		washedChannel := Channel{
			ID:                fmt.Sprintf("%d", channel.ID),
			SortOrder:         channel.Title,
			ChannelName:       channel.Title,
			ChannelType:       channel.Title,
			ChannelCode:       channel.Title,
			GroupID:           channel.GroupID,
			RelevantUnits:     []any{},
			Mandatory:         channel.Mandatory,
			RelevantCountires: []any{},
			RegionID:          channel.RegionID,
			NewsCategoryID:    channel.NewsCategoryID,
		}

		for _, relevantCountry := range channel.RelevantCountires {
			washedChannel.RelevantCountires = append(washedChannel.RelevantCountires, relevantCountry)
		}
		for _, relevantUnit := range channel.RelevantUnits {
			washedChannel.RelevantUnits = append(washedChannel.RelevantUnits, relevantUnit)
		}

		washedData.Channels = append(washedData.Channels, washedChannel)
	}
	return washedData, nil

}
func Wash(inputfilename string, outputfilename string) error {
	profileData := schemas.Profiledata{}
	data, err := os.ReadFile(inputfilename)

	if err != nil {
		return err
	}

	json.Unmarshal(data, &profileData)

	washed, err := WashProfileData(&profileData)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(washed, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(outputfilename, output, 0644)
	if err != nil {
		return err
	}
	return nil
}
