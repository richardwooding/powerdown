package model

import "time"

type AreaResponse struct {
	Events []struct {
		End   time.Time `json:"end"`
		Note  string    `json:"note"`
		Start time.Time `json:"start"`
	} `json:"events"`
	Info struct {
		Name   string `json:"name"`
		Region string `json:"region"`
	} `json:"info"`
	Schedule struct {
		Days []struct {
			Date   string     `json:"date"`
			Name   string     `json:"name"`
			Stages [][]string `json:"stages"`
		} `json:"days"`
		Source string `json:"source"`
	} `json:"schedule"`
}
