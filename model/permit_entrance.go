package model

type PermitEntrance struct {
	//Permit Entrance ID
	PermitEntranceID string

	//Facility ID the permit belongs to
	FacilityID string

	PermitEntranceName string
	//Permit Entrance name

	//Permit Entrance description
	PermitEntranceDescription string

	//District the permit resides in
	District string

	//Town the permit resides in
	Town string

	//Is the permit accessible by vehicle
	PermitEntranceAccessible bool

	//Latitude of the permit location
	Longitude float64

	//Longitude of the permit location
	Latitude float64

	GEOSJON GeoJson

	//Record creation date
	CreatedDate string

	//Record last update date
	LastUpdatedDate string

	ATTRIBUTES []Attribute
	//ENTITYMEDIA*	[...]
	//ZONES []Zones
}
