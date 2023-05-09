package models

import (
	"github.com/jinzhu/gorm"
)

type Fence struct {
	gorm.Model
	Name         string     `json:"name"`
	Locations    []Location `gorm:"ForeignKey:FenceID" json:"locations"`
	IncludeBikes []Bike     `gorm:"many2many:fence_bikes;"`
}

// IsPointInside 判断给定的经纬度是否在电子围栏内
func (f *Fence) IsPointInside(lat, lon float64) bool {
	point := Location{Latitude: lat, Longitude: lon}
	n := len(f.Locations)
	inside := false

	j := n - 1
	for i := 0; i < n; i++ {
		if (f.Locations[i].Latitude > point.Latitude) != (f.Locations[j].Latitude > point.Latitude) &&
			(point.Longitude < (f.Locations[j].Longitude-f.Locations[i].Longitude)*(point.Latitude-f.Locations[i].Latitude)/
				(f.Locations[j].Latitude-f.Locations[i].Latitude)+f.Locations[i].Longitude) {
			inside = !inside
		}
		j = i
	}

	return inside
}
