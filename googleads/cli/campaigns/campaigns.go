package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	gads "github.com/denton/gads/googleads"
)

var configJson = flag.String("oauth", "./oauth.json", "API credentials")

func main() {
	config, err := gads.NewCredentialsFromFile(*configJson)
	if err != nil {
		log.Fatal(err)
	}

	var pageSize int64 = 500
	var offset int64

	// show all Campaigns
	cs := gads.NewCampaignService(&config.Auth)
	paging := gads.Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	fmt.Printf("\nCampaigns\n")
	for {
		foundCampaigns, totalCount, err := cs.Get(
			gads.Selector{
				Fields: []string{
					"Id",
					"Name",
					"Status",
					"ServingStatus",
					"StartDate",
					"EndDate",
					"Settings",
					"AdvertisingChannelType",
					"AdvertisingChannelSubType",
					"Labels",
					"TrackingUrlTemplate",
					"UrlCustomParameters",
				},
				Predicates: []gads.Predicate{
					{"Status", "EQUALS", []string{"PAUSED"}},
				},
				Ordering: []gads.OrderBy{
					{"Id", "ASCENDING"},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, campaign := range foundCampaigns {
			campaignJson, _ := json.MarshalIndent(campaign, "", "  ")
			fmt.Printf("%s\n", campaignJson)
		}
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}

}
