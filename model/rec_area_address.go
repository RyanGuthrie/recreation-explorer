package model

type RecreationAreaAddress struct {
	// RecArea Address ID
	RecAreaAddressID string

	// RecArea ID
	RecAreaID string

	// Address type for the RecArea
	RecAreaAddressType string

	// Address line 1 of the RecArea
	RecAreaStreetAddress1 string

	// Address line 2 of the RecArea
	RecAreaStreetAddress2 string

	// Address line 3 of the RecArea
	RecAreaStreetAddress3 string

	// City where the RecArea is located
	City string

	// Postal code for the RecArea
	PostalCode string

	// State code for the RecArea
	AddressStateCode string

	// Abbreviated country code for the RecArea
	AddressCountryCode string

	// Record last update date
	LastUpdatedDate string
}
