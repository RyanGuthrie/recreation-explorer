package model

type Link struct {
	//Primary Key
	EntityLinkID string

	//Type of link, e.g.Facebook, Twitter, official site
	LinkType string

	//RecArea ID OR Facility ID
	EntityID string

	//RecArea or Facility
	EntityType string

	//Full text title
	Title string

	//Text description of the entity link
	Description string

	//Internet address (URL) to a web site
	URL string
}
