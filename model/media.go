package model

type Media struct {
	//Type of Media, e.g. Image, Video, etc.
	MediaType string

	//Primary Key
	EntityMediaID string

	//RecArea ID OR Facility ID OR Tour ID OR Permit Entrance ID OR Campsite ID
	EntityID string

	//RecArea, Facility, Tour, Entrance, or Site
	EntityType string

	//Full title of the entity media
	Title string

	//Optional subtitle of the entity media
	Subtitle string

	//Optional description of the entity media
	Description string

	//Optional embedded code for media entity
	EmbedCode string

	//Height in pixels for media image
	Height int

	//Width in pixels for the media image
	Width int

	//Whether the image is a primary image
	IsPrimary bool

	//Whether the image is a preview image
	IsPreview bool

	//Whether the image is a gallery image
	IsGallery bool

	//Internet address (URL) to the entity media
	URL string

	//Optional credit for entity media
	Credits string
}
