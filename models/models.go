package models

import "time"

type BookDetails struct {
	URL           string   `json:"url"`
	Name          string   `json:"name"`
	Isbn          string   `json:"isbn"`
	Authors       []string `json:"authors"`
	NumberOfPages int      `json:"numberOfPages"`
	Publisher     string   `json:"publisher"`
	Country       string   `json:"country"`
	MediaType     string   `json:"mediaType"`
	Released      string   `json:"released"`
	Characters    []string `json:"characters"`
	PovCharacters []string `json:"povCharacters"`
}

type Comments struct {
	ID        string     `json:"id"`
	Book      string     `json:"book"`
	Comment   string     `json:"comment"`
	IPAddress string     `json:"ip"`
	CreatedAt *time.Time `json:"created_at"`
}
