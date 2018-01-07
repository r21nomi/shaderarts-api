package datastore

type Program struct {
	ID string `json:"id"`
	Type int `json:"type"`
	Code string `json:"code"`
	ArtID string `json:"artId"`
}