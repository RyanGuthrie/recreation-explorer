package domain

import "fmt"

type Gaia struct {
	Latitude  float64
	Longitude float64
}

func (g Gaia) Link() string {
	return fmt.Sprintf(
		"https://www.gaiagps.com/map/?loc=17.5/%f/%f&layer=mineral-resources,usfs-mvum,usfs-roads,usfs-recsites,GaiaOverlandRasterFeet",
		g.Longitude, g.Latitude)
}
