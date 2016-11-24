package mgeoloc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const freegeoip = "https://freegeoip.net/json/"

type GeographicLocation struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zipcode"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitutde"`
}

func FromAddr(addr string) GeographicLocation {
	var resp *http.Response
	var err error
	var body []byte

	if resp, err = http.Get(freegeoip + addr); err != nil {
		log.Printf("failed to get location: %s", err)
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("failed to read response body: %s", err)
	}

	var geoloc GeographicLocation

	if err = json.Unmarshal(body, &geoloc); err != nil {
		log.Printf("failed to unmarshal response body: %s", err)
	}

	return geoloc
}

func (geo *GeographicLocation) FormatData() string {
	return fmt.Sprintf("[ IP: %s ]\n[ Country Code: %s ]\n[ Country Name: %s ]\n[ City: %s ]\n[ Zip Code: %s ]\n[ Area Code: %d ]\n[ Metro Code: %d ]\n[ Lat: %f ]\n[ Long: %f ]\n",
		geo.IP,
		geo.CountryCode,
		geo.CountryName,
		geo.City,
		geo.ZipCode,
		geo.AreaCode,
		geo.MetroCode,
		geo.Latitude,
		geo.Longitude,
	)
}
