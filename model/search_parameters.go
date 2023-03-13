package model

import "fmt"

type SearchParameters struct {
	Query  string `json:"query"`
	Limit  int    `json:"LIMIT"`
	Offset int    `json:"OFFSET"`
}

func (sp SearchParameters) String() string {
	return fmt.Sprintf("Query [%s], Limit [%d], Offset [%d]", sp.Query, sp.Limit, sp.Offset)
}
