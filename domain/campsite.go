package domain

import (
	"fmt"
	"log"
	"server/model"
	"strings"
)

type Campsite struct {
	FacilityID string
	CampsiteID string
	Campsite   model.Campsite
}

func (c Campsite) Explore() error {
	mc := c.Campsite

	var attributesSummary []string
	for _, attribute := range mc.Attributes {
		attributesSummary = append(attributesSummary, fmt.Sprintf("%s: %s", attribute.AttributeName, attribute.AttributeValue))
	}
	var permittedEquipmentSummary []string
	for _, eq := range mc.PermittedEquipment {
		permittedEquipmentSummary = append(permittedEquipmentSummary, fmt.Sprintf("%s (length: %v)", eq.EquipmentName, eq.MaxLength))
	}

	log.Printf(fmt.Sprintf(""+
		"Name: %s, Type: %s, Use: %s\n"+
		"Loop                %s\n"+
		"Vehicle Accessible: %v\n"+
		"Reservable:         %v\n"+
		"Coordinates:        (%f, %f) %s\n"+
		"Attributes:\n-  %s\n"+
		"Permitted Equipment:\n-  %s\n",
		mc.CampsiteName, mc.CampsiteType, mc.TypeOfUse,
		mc.Loop,
		mc.CampsiteAccessible,
		mc.CampsiteReservable,
		mc.CampsiteLatitude, mc.CampsiteLongitude, Gaia{mc.CampsiteLatitude, mc.CampsiteLongitude}.Link(),
		strings.Join(attributesSummary, "\n-  "),
		strings.Join(permittedEquipmentSummary, "\n-  "),
	))

	return nil
}
