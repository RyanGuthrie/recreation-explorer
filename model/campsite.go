package model

import "fmt"

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

	//Is the campsite reservable
	CampsiteReservable bool

	//Longitude of the permit
	CampsiteLongitude float64

	//Latitude of the permit
	CampsiteLatitude float64

	//Record creation date
	CreatedDate string

	//Record last update date
	LastUpdatedDate string

	Attributes         []Attribute          `json:"ATTRIBUTES"`
	PermittedEquipment []PermittedEquipment `json:"PERMITTEDEQUIPMENT"`
	EntityMedia        []Media              `json:"ENTITYMEDIA"`
}

func (c Campsite) String() string {
	nameSummary := fmt.Sprintf("%-35s %10s", c.CampsiteName, c.CampsiteID)
	return fmt.Sprintf("Name [%s], Facility [%-10s] Reservable: [%-5v], TypeOfUse [%-9s], Type [%-30s]", nameSummary, c.FacilityID, c.CampsiteReservable, c.TypeOfUse, c.CampsiteType)
}
