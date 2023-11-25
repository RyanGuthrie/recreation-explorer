package model

import "fmt"

type Facility struct {
	//RIDB unique Facility ID
	FacilityID string

	//Legacy Facility ID
	LegacyFacilityID string

	//The agency's internal Facility ID provided to the RIDB by the agency
	OrgFacilityID string

	//The parent Organization ID
	ParentOrgID string

	//The parent RecArea ID
	ParentRecAreaID string

	//Full name of the Facility
	FacilityName string

	//Text describing the main features of the Facility
	FacilityDescription string

	//Description of the type of Facility
	FacilityTypeDescription string

	//Text describing monetary charges associated with entrance to or usage of the Facility
	FacilityUseFeeDescription string

	//Text that provides general directions and/or the general location of the Facility
	FacilityDirections string

	//Phone number of the Facility
	FacilityPhone string

	//Email address of the Facility
	FacilityEmail string

	//Internet address (URL) for the web site hosting the reservation system
	FacilityReservationURL string

	//Internet address (URL) that hosts the Facility map
	FacilityMapURL string

	//Information about the Americans with Disabilities Act accessibility for the Facility
	FacilityAdaAccess string

	//
	FacilityAccessibilityText string

	GeoJSON GeoJson `json:"GEOJSON"`

	//Longitude in decimal degrees -180.0 to 180.0
	FacilityLongitude float64

	//Latitude in decimal degrees -90.0 to 90.0
	FacilityLatitude float64

	//Details on the stay limits for the Facility
	StayLimit string

	//List of keywords for the Facility
	Keywords string

	//Whether the Facility is reservable
	Reservable bool

	//Whether the Facility is enabled
	Enabled bool

	//Record last update date
	LastUpdatedDate string

	Campsites         []Campsite         `json:"CAMPSITE"`
	PermitEntrances   []PermitEntrance   `json:"PERMITENTRANCE"`
	Organizations     []Organization     `json:"ORGANIZATION"`
	RecAreas          []RecreationArea   `json:"RECAREA"`
	FacilityAddresses []FacilityAddress  `json:"FACILITYADDRESS"`
	Activities        []FacilityActivity `json:"ACTIVITY"`
	Events            []Event            `json:"EVENT"`
	Links             []Link             `json:"LINK"`
	Media             []Media            `json:"MEDIA"`
	Tours             []FacilityTour     `json:"TOUR"`
}

func (f *Facility) String() string {
	return fmt.Sprintf("%-45s ID [%s], RecAreas(#%d): %v", f.FacilityName, f.FacilityID, len(f.RecAreas), f.RecAreas)
}
