package model

type Organization struct {
	//Organization ID
	OrgID string

	//Full name of organization
	OrgName string

	//Internet address (URL) that hosts the sample image or photo of the organization
	OrgImageURL string

	//Optional Readable text that provides the URL address link
	OrgURLText string

	//Internet address (URL) for web site of the organization responsible for submitting and maintaining the data to be exchanged
	OrgURLAddress string

	//Internet address (URL) that hosts the sample image or photo of the organization
	OrgType string

	//Abbreviated name of the organization
	OrgAbbrevName string

	//Organization jurisdiction type
	OrgJurisdictionType string

	//Parent Organization ID
	OrgParentID string

	//Record last update date
	LastUpdatedDate string
}
