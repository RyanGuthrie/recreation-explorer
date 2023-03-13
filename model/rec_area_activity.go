package model

type RecreationAreaActivity struct {
	//Activity ID
	ActivityID int

	//Parent ID of the related Activity
	ActivityParentID string

	//Parent RecArea ID of the Activity
	RecAreaID string

	//Name of the Activity
	ActivityName string

	//Description of the Activity
	RecAreaActivityDescription string

	//Text describing monetary charges associated with the Activity
	RecAreaActivityFeeDescription string
}
