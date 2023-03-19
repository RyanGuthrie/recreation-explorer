package model

import "fmt"

type RecreationArea struct {
	//RIDB unique RecArea ID
	RecAreaID string

	//The agency's internal RecArea ID provided to the RIDB by the agency
	OrgRecAreaID string

	//RecArea link
	ResourceLink string

	//Parent Organization ID
	ParentOrgID string

	//Full name of the RecArea
	RecAreaName string

	//Text that describes the RecArea
	RecAreaDescription string

	//Text describing monetary charges associated with entrance to or usage of a RecArea
	RecAreaFeeDescription string

	//Directions to the RecArea
	RecAreaDirections string

	//Phone number for RecArea
	RecAreaPhone string

	//Email address of the RecArea
	RecAreaEmail string

	//Internet address (URL) for the web site hosting the reservation system
	RecAreaReservationURL string

	//Internet address (URL) that hosts the RecArea map
	RecAreaMapURL string

	GEOJSON GeoJson

	//Longitude in decimal degrees -180.0 to 180.0
	RecAreaLongitude float64

	//Latitude in decimal degrees -90.0 to 90.0
	RecAreaLatitude float64

	//Details on the stay limits for the RecArea
	StayLimit string

	//List of keywords for the RecArea
	Keywords string

	//Whether the RecArea is reservable
	Reservable bool

	//Whether the RecArea is enabled
	Enabled bool

	//Record last update date
	LastUpdatedDate string

	Organizations    []Organization           `json:"ORGANIZATION"`
	Facilities       []RecreationAreaFacility `json:"FACILITY"`
	RecAreaAddresses []RecreationAreaAddress  `json:"RECAREAADDRESS"`
	Activities       []RecreationAreaActivity `json:"ACTIVITY"`
	Events           []Event                  `json:"EVENT"`
	Media            []Media                  `json:"MEDIA"`
	Links            []Link                   `json:"LINK"`
}

func (ra RecreationArea) String() string {
	return fmt.Sprintf("Name: %s, Organizations(%d), Facilities (%d), Activities (%d), Events (%d)",
		ra.RecAreaName,
		len(ra.Organizations),
		len(ra.Facilities),
		len(ra.Activities),
		len(ra.Events),
	)
}
