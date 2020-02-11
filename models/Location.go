package models

import "gopkg.in/guregu/null.v3"

type Location struct {
	Id             string      `json:"id"`
	OrganizationId string      `json:"organization_id"`
	Name           string      `json:"location_name,string"`
	Street1        null.String `json:"location_street_1,string"`
	Street2        null.String `json:"location_street_2,string"`
	City           null.String `json:"location_city,string"`
	State          null.String `json:"location_state,string"`
	ZipCode        null.String `json:"location_zip_code,string"`
	GeoPoint       null.String `json:"geo_point"`
	Latitude       null.String `json:"location_latitude,string"`
	Longitude      null.String `json:"location_longitude,string"`
}
