package datastore

type Program struct {
	ID string `json:"id"`
	Type int `json:"type"`
	Code string `gorm:"type:text" json:"code"`
	ArtID string `json:"artId"`
}