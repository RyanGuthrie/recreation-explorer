package model

import "fmt"

type Results struct {
	CurrentCount int `json:"CURRENT_COUNT"`
	TotalCount   int `json:"TOTAL_COUNT"`
}

func (r Results) String() string {
	return fmt.Sprintf("Counts: Current [%d], Total [%d]", r.CurrentCount, r.TotalCount)
}
