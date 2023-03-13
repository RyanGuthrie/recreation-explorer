package model

type FacilityActivity struct {
	//Activity ID
	ActivityID string

	//Parent Facility ID of the Activity
	FacilityID string

	//Name of the Activity
	ActivityName string

	//Description of the Activity
	FacilityActivityDescription string

	//Text describing monetary charges associated with the Activity
	FacilityActivityFeeDescription string
}
