package model

import "fmt"

type Metadata struct {
	Results          Results          `json:"RESULTS"`
	SearchParameters SearchParameters `json:"SEARCH_PARAMETERS"`
}

func (md Metadata) String() string {
	return fmt.Sprintf("Results: %v, SearchParameters: %v", md.Results, md.SearchParameters)
}
