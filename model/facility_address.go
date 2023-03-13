package model

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
