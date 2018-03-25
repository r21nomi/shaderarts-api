package datastore

type Code struct {
	ID    string
	Type  int
	Text  string `gorm:"type:text"`
	ArtID string
}
