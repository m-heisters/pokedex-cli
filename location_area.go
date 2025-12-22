package main

type location_areas_response struct {
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
}
