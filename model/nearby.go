package model

type NearbyResponse struct {
	Areas []struct {
		Count int    `json:"count"`
		Id    string `json:"id"`
	} `json:"areas"`
}