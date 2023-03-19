package model

import (
	"fmt"
	"strings"
)

type FacilityAddress struct {
	//Facility Address ID
	FacilityAddressID string

	//Facility ID
	FacilityID string

	//Address type for the Facility
	FacilityAddressType string

	//Address line 1 of the Facility
	FacilityStreetAddress1 string

	//Address line 2 of the Facility
	FacilityStreetAddress2 string

	//Address line 3 of the Facility
	FacilityStreetAddress3 string

	//City where the Facility is located
	City string

	//Postal code for the Facility
	PostalCode string

	//State code for the Facility
	AddressStateCode string

	//Abbreviated country code for the Facility Address
	AddressCountryCode string

	//Record last update date
	LastUpdatedDate string
}

func (address FacilityAddress) String() string {
	var summary []string
	if address.FacilityStreetAddress1 != "" {
		summary = append(summary, address.FacilityStreetAddress1)
	}
	if address.FacilityStreetAddress2 != "" {
		summary = append(summary, address.FacilityStreetAddress2)
	}
	if address.FacilityStreetAddress3 != "" {
		summary = append(summary, address.FacilityStreetAddress3)
	}

	if address.City != "" {
		summary = append(summary, fmt.Sprintf("%s, %s %s",
			address.City,
			address.AddressStateCode,
			address.PostalCode))

	} else if address.AddressStateCode != "" {
		summary = append(summary, address.AddressStateCode)
	}

	return strings.Join(summary, "\n")
}
