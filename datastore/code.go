package datastore

type Code struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Text  string `gorm:"type:text" json:"text"`
	ArtID string `json:"artId"`
}
