package model

import "fmt"

type Reservation struct {
	HistoricalReservationID string
	OrderNumber             string
	Agency                  string
	OrgID                   string
	CodeHierarchy           string
	RegionCode              string
	RegionDescription       string
	ParentLocationID        string
	ParentLocation          string
	LegacyFacilityID        string
	Park                    string
	SiteType                string
	UseType                 string
	ProductID               string
	InventoryType           string
	FacilityID              string
	FacilityZIP             string
	FacilityState           string
	FacilityLongitude       string
	FacilityLatitude        string
	CustomerZIP             string
	Tax                     string
	UseFee                  string
	TranFee                 string
	AttrFee                 string
	TotalBeforeTax          string
	Discount                string
	TotalPaid               string
	StartDate               string
	EndDate                 string
	OrderDate               string
	Nights                  string
	NumberOfPeople          string
	EquipmentDescription    string
	EquipmentLength         string
}

func (r Reservation) String() string {
	return fmt.Sprintf("Order #[%v] Date [%v - %v] Paid: [%v]", r.OrderNumber, r.StartDate, r.EndDate, r.TotalPaid)
}
