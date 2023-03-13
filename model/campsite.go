package model

type Campsite struct {
	//Campsite ID
	CampsiteID string

	//Facility ID the campsite belongs to
	FacilityID string

	//Campsite name
	CampsiteName string

	//Campsite type
	CampsiteType string

	//Type of use
	TypeOfUse string

	//Name of loop the campsite resides on
	Loop string

	//Is the campsite accessible by vehicle
	CampsiteAccessible bool

	//Longitude of the permit
	CampsiteLongitude float64

	//Latitude of the permit
	CampsiteLatitude float64

	//Record creation date
	CreatedDate string

	//Record last update date
	LastUpdatedDate string

	ATTRIBUTES         []Attribute
	PERMITTEDEQUIPMENT []PermittedEquipment

	//ENTITYMEDIA*	[...]
}
