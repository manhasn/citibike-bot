package main

type stationStatus struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []struct {
			StationID             string   `json:"station_id"`
			Name                  string   `json:"name"`
			ShortName             string   `json:"short_name"`
			Lat                   float64  `json:"lat"`
			Lon                   float64  `json:"lon"`
			RegionID              int      `json:"region_id,omitempty"`
			RentalMethods         []string `json:"rental_methods"`
			Capacity              int      `json:"capacity"`
			RentalURL             string   `json:"rental_url"`
			EightdHasKeyDispenser bool     `json:"eightd_has_key_dispenser"`
			EightdStationServices []struct {
				ID                  string      `json:"id"`
				ServiceType         string      `json:"service_type"`
				BikesAvailability   string      `json:"bikes_availability"`
				DocksAvailability   string      `json:"docks_availability"`
				Name                interface{} `json:"name"`
				Description         interface{} `json:"description"`
				ScheduleDescription interface{} `json:"schedule_description"`
				LinkForMoreInfo     interface{} `json:"link_for_more_info"`
			} `json:"eightd_station_services,omitempty"`
		} `json:"stations"`
	} `json:"data"`
}
