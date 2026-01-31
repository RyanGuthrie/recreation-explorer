package model

import "fmt"

type Facility struct {
	//RIDB unique Facility ID
	FacilityID string `json:"FacilityID"`

	//Legacy Facility ID
	LegacyFacilityID string `json:"LegacyFacilityID"`

	//The agency's internal Facility ID provided to the RIDB by the agency
	OrgFacilityID string `json:"OrgFacilityID"`

	//The parent Organization ID
	ParentOrgID string `json:"ParentOrgID"`

	//The parent RecArea ID
	ParentRecAreaID string `json:"ParentRecAreaID"`

	//Full name of the Facility
	FacilityName string `json:"FacilityName"`

	//Text describing the main features of the Facility
	FacilityDescription string `json:"FacilityDescription"`

	//Description of the type of Facility
	FacilityTypeDescription string `json:"FacilityTypeDescription"`

	//Text describing monetary charges associated with entrance to or usage of the Facility
	FacilityUseFeeDescription string `json:"FacilityUseFeeDescription"`

	//Text that provides general directions and/or the general location of the Facility
	FacilityDirections string `json:"FacilityDirections"`

	//Phone number of the Facility
	FacilityPhone string `json:"FacilityPhone"`

	//Email address of the Facility
	FacilityEmail string `json:"FacilityEmail"`

	//Internet address (URL) for the web site hosting the reservation system
	FacilityReservationURL string `json:"FacilityReservationURL"`

	//Internet address (URL) that hosts the Facility map
	FacilityMapURL string `json:"FacilityMapURL"`

	//Information about the Americans with Disabilities Act accessibility for the Facility
	FacilityAdaAccess string `json:"FacilityAdaAccess"`

	//
	FacilityAccessibilityText string `json:"FacilityAccessibilityText"`

	GeoJSON GeoJson `json:"GEOJSON"`

	//Longitude in decimal degrees -180.0 to 180.0
	FacilityLongitude float64 `json:"FacilityLongitude"`

	//Latitude in decimal degrees -90.0 to 90.0
	FacilityLatitude float64 `json:"FacilityLatitude"`

	//Details on the stay limits for the Facility
	StayLimit string `json:"StayLimit"`

	//List of keywords for the Facility
	Keywords string `json:"Keywords"`

	//Whether the Facility is reservable
	Reservable bool `json:"Reservable"`

	//Whether the Facility is enabled
	Enabled bool `json:"Enabled"`

	//Record last update date
	LastUpdatedDate string `json:"LastUpdatedDate"`

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
