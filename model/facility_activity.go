package model

import (
	"fmt"
	"strings"
)

type FacilityActivity struct {
	//Activity ID
	ActivityID int

	//Parent Facility ID of the Activity
	FacilityID string

	//Name of the Activity
	ActivityName string

	//Description of the Activity
	FacilityActivityDescription string

	//Text describing monetary charges associated with the Activity
	FacilityActivityFeeDescription string
}

func (activity FacilityActivity) String() string {
	var summary []string

	summary = append(summary, fmt.Sprintf("%s (%d)", activity.ActivityName, activity.ActivityID))

	if activity.FacilityActivityFeeDescription != "" {
		summary = append(summary, fmt.Sprintf("Fee: %s", activity.FacilityActivityFeeDescription))
	}

	summary = append(summary, activity.FacilityActivityDescription)

	return fmt.Sprintf(strings.Join(summary, "\n  "))
}
