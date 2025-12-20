package main

type location_areas struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
}
