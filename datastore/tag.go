package datastore

type Tag struct {
	ID   string
	Text string `gorm:"type:text"`
}
